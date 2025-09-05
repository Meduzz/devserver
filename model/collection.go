package model

import (
	"encoding/json"
)

/*
TODO
* Collection, flags for read/write/delete permissions?
* More kinds of fields (file?, md?, relation?, decimal?, id (with id-strategy)?, any?, oneof/condition/routing?)
*/

type (
	CKind    string
	FKind    string
	Strategy string

	Collection struct {
		Kind   CKind           `json:"kind"`
		Name   string          `json:"name"`
		Label  string          `json:"label"`
		Config json.RawMessage `json:"config"` // *File|*Api|*Static
		Fields []*Field        `json:"fields"`
	}

	Field struct {
		Kind     FKind    `json:"kind"`
		Name     string   `json:"name"`
		Label    string   `json:"label"`
		Default  any      `json:"any,omitempty"`
		Strategy Strategy `json:"strategy,omitempty"`
		Optional bool     `json:"optional"`
		Fields   []*Field `json:"field,omitempty"`
	}

	// json file as a db
	FileCollection struct {
		File string `json:"file"`
		ID   string `json:"id"` // id column
	}

	// proxy requests
	ApiCollection struct {
		Host string `json:"host"`
	}

	// configuration (in handy.yaml) as db
	StaticCollection struct {
		Key string `json:"key"` // key on app app.<key> = [<collection.data>]
	}
)

const (
	FileCollectionKind   = CKind("file")   // map requests to a file
	ApiCollectionKind    = CKind("api")    // forward request to api
	StaticCollectionKind = CKind("static") // this is a collection that cant be mutated outside cms (think settings)
)

const (
	String   = FKind("string") // input[type=text]
	Text     = FKind("text")   // textarea
	Number   = FKind("number")
	Boolean  = FKind("boolean")
	Date     = FKind("date")
	Datetime = FKind("datetime")
	Select   = FKind("select") // select|radio
	List     = FKind("list")   // checkbox|tags|object?
	Colour   = FKind("colour")
	Object   = FKind("object")
)

const (
	Dropdown = Strategy("dropdown")
	Radio    = Strategy("radio")
	Checkbox = Strategy("checkbox")
	Tags     = Strategy("tags")
)

var (
	CollectionTypes = map[CKind]string{}
	FieldTypes      = map[FKind]string{}
)

func init() {
	CollectionTypes[FileCollectionKind] = ""
	CollectionTypes[ApiCollectionKind] = ""

	FieldTypes[String] = ""
	FieldTypes[Text] = ""
	FieldTypes[Number] = ""
	FieldTypes[Boolean] = ""
	FieldTypes[Date] = ""
	FieldTypes[Datetime] = ""
	FieldTypes[Select] = ""
	FieldTypes[List] = ""
	FieldTypes[Colour] = ""
}
