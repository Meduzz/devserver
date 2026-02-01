package model

import "encoding/json"

type (
	EndpointKind string

	// TODO some endpoints will need to be able to tell how they need to be exposed (later).
	Endpoint struct {
		Kind   EndpointKind    `json:"kind"`
		Name   string          `json:"name"`
		Path   string          `json:"path"`
		Drop   string          `json:"drop,omitempty"` // drop prefix
		Config json.RawMessage `json:"config"`         // *Static|*Proxy
	}

	// TODO allow exposing the content of a zip-file as static resource
	StaticEndpoint struct {
		Dir string `json:"dir"`
		SPA string `json:"spa"` // index file
	}

	ProxyEndpoint struct {
		Host string `json:"host"`
	}
)

const (
	StaticEndpointKind = EndpointKind("static") // {static.path}/ -> /{fs.path}
	ProxyEndpointKind  = EndpointKind("proxy")  // /{path} -> {proxy.host}
)
