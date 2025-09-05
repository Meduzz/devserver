package filedb

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	fileDb struct {
		id     string
		idType model.FKind
		data   []map[string]any
		file   string
	}
)

func NewFileDB(file, id string, fields []*model.Field) (DB, error) {
	idCol := slice.Head(slice.Filter(fields, func(f *model.Field) bool {
		return f.Name == id
	}))

	if idCol == nil {
		return nil, fmt.Errorf("no field named %s (id) was present", id)
	}

	bs, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	data := make([]map[string]any, 0)

	err = json.Unmarshal(bs, &data)

	if err != nil {
		return nil, err
	}

	return &fileDb{
		id:     id,
		idType: idCol.Kind,
		data:   data,
		file:   file,
	}, nil
}

// Create implements DB.
func (f *fileDb) Create(data map[string]any) (map[string]any, error) {
	f.data = append(f.data, data)

	return data, nil
}

// Delete implements DB.
func (f *fileDb) Delete(id string) error {
	switch f.idType {
	case model.Number:
		it, err := strconv.Atoi(id)

		if err != nil {
			return err
		}

		f.data = slice.Filter(f.data, func(m map[string]any) bool {
			return m[f.id] != it
		})
	default:
		f.data = slice.Filter(f.data, func(m map[string]any) bool {
			return m[f.id] != id
		})
	}

	return nil
}

// List implements DB.
func (f *fileDb) List() ([]map[string]any, error) {
	return f.data, nil
}

// Update implements DB.
func (f *fileDb) Update(id string, data map[string]any) (map[string]any, error) {
	switch f.idType {
	case model.Number:
		it, err := strconv.Atoi(id)

		if err != nil {
			return nil, err
		}

		f.data = slice.Map(f.data, func(m map[string]any) map[string]any {
			if m[f.id] != it {
				return m
			} else {
				return data
			}
		})
	default:
		f.data = slice.Map(f.data, func(m map[string]any) map[string]any {
			if m[f.id] != id {
				return m
			} else {
				return data
			}
		})
	}

	return data, nil
}

func (f *fileDb) Save() error {
	bs, err := json.Marshal(f.data)

	if err != nil {
		return err
	}

	return os.WriteFile(f.file, bs, 0x755)
}
