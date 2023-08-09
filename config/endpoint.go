package config

type (
	// Define an endpoint, that can get mapped to an path
	Endpoint struct {
		Kind    string   `json:"kind"` // proxy|api|command|static
		Method  string   `json:"method"`
		Path    string   `json:"path"`
		Proxy   *Proxy   `json:"proxy,omitempty"`
		Api     *Api     `json:"api,omitempty"`
		Command *Command `json:"command,omitempty"`
		Static  *Static  `json:"static,omitempty"`
	}

	// Define a proxy to be bound to a path.
	Proxy struct {
		Kind string     `json:"kind"`
		Http *HttpProxy `json:"http"` // http
	}

	// Define the proxy as a http proxy.
	HttpProxy struct {
		Target string `json:"target"`
	}

	// TODO wendy proxy?

	// Define the endpoint as an api endpoint.
	Api struct {
		Datasource string `json:"datasource"`
		// TODO mapping from path & query params?
	}

	// Define a command to be executed on an endpoint.
	Command struct {
		Command string `json:"command"`
		// TODO mapping from path & query params?
	}

	// Define some static content to use in an endpoint
	Static struct {
		Datasource string `json:"datasource"`
	}

	// Define template(s) as content to use in an endpoint
	Template struct {
		Kind       string `json:"kind"` // mustache|go
		Datasource string `json:"datasource"`
		// TODO ContentType always html? Not sure.
	}

	// TODO webassembly endpoint would be cool...
)
