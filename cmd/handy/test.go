package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Meduzz/commando"
	com "github.com/Meduzz/commando/model"
	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
)

func init() {
	t := commando.Command("test", func(ec com.ExecuteCommand) error {
		col := &model.Collection{
			Fields: []*model.Field{
				{
					Name:  "name",
					Label: "Name",
					Kind:  model.String,
				}, {
					Name:  "age",
					Label: "Age",
					Kind:  model.Number,
				}, {
					Name:     "greeting",
					Label:    "Greeting",
					Kind:     model.Object,
					Optional: true,
					Fields: []*model.Field{
						{
							Name:  "text",
							Label: "Text",
							Kind:  model.String,
						},
					},
				},
			},
		}

		fields := slice.Map(col.Fields, toField)

		typ := reflect.StructOf(fields)
		it := reflect.New(typ).Interface()
		data := []byte(`{"name":"Test", "age":10, "greeting":{"text":"Hello %s"}}`)
		err := json.Unmarshal(data, it)

		if err != nil {
			return err
		}

		fmt.Printf("%v\n", it)

		return nil
	})

	t.Description = "Testing stuff"
}

func toField(f *model.Field) reflect.StructField {
	tags := []string{f.Name}

	if f.Optional {
		tags = append(tags, "omitempty")
	}

	t := fmt.Sprintf(`json:"%s"`, strings.Join(tags, ",s"))

	return reflect.StructField{
		Name: strings.Title(f.Name),
		Type: fieldType(f),
		Tag:  reflect.StructTag(t),
	}
}

func fieldType(f *model.Field) reflect.Type {
	switch f.Kind {
	case model.String, model.Text, model.Date, model.Datetime, model.Colour:
		return reflect.TypeOf("")
	case model.Boolean:
		return reflect.TypeOf(false)
	case model.Number:
		return reflect.TypeOf(0)
	case model.List:
		return reflect.ArrayOf(0, fieldType(f.Fields[0])) // TODO not safe, but somewhat safe
	case model.Object:
		fields := slice.Map(f.Fields, toField)
		return reflect.StructOf(fields)
	}
	return nil
}
