package config

import "encoding/json"

type (
	Config struct {
		App         *App          `json:"app"`
		Endpoints   []*Endpoint   `json:"endpoints"`
		Datasources []*Datasource `json:"datasources,omitempty"`
		json.RawMessage
	}

	App struct {
		Name   string `json:"name"`
		Prefix string `json:"prefix"` // used when starting with a dir?
	}
)
