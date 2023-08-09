package config

type (
	// Define a datasource to be used by an API endpoint.
	Datasource struct {
		Kind      string     `json:"kind"`     // file|property|directory
		Encoding  string     `json:"encoding"` // <empty>|json|html|js|css|<etc>
		Name      string     `json:"name"`
		File      *File      `json:"file,omitempty"`
		Property  *Property  `json:"property,omitempty"`
		Directory *Directory `json:"directory,omitempty"`
		// TODO optional schema?
		// TODO read/write settings?
	}

	// Use data in a file available at the url as datasource
	File struct {
		URL string `json:"url"`
		// TODO secrets and stuff for http urls
	}

	// Use a field in the config as datasource.
	Property struct {
		Field   string `json:"field"`
		Base64  bool   `json:"base64"`
		Kind    string `json:"kind"`    // object|list|keyed
		IdField string `json:"idfield"` // query param for updates/deletes (also in stored data)
	}

	// Use a directory as a datasource
	Directory struct {
		Dir    string `json:"dir"`
		Filter string `json:"filter"` // glob? file ending?
	}

	// TODO sqlite, sql, badger, bolt etc
)
