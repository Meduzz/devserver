package model

/*
 App is the config of the app that is used by the various commands.
 An app can be loaded either by file, or by url.
*/

type (
	App struct {
		Name      string            `json:"name"`
		Endpoints []*Endpoint       `json:"endpoints,omitempty"`
		Children  map[string]string `json:"children,omitempty"` // key -> relative path | url
	}
)
