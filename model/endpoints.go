package model

import "encoding/json"

/*
TODO
* Static, flag to allow browsing static files?
* Proxy, settings to drop/add path?
* Page, setting for layout?
*/

type (
	EKind string

	Endpoint struct {
		Kind   EKind           `json:"kind"`
		Name   string          `json:"name"`
		Path   string          `json:"path"`
		Config json.RawMessage `json:"config"` // *Static|*Collection|*Proxy|*Page
	}

	StaticEndpoint struct {
		Dir       string `json:"dir"`
		SPA       *bool  `json:"spa"`
		Browsable *bool  `json:"browsable"`
	}

	CollectionEndpoint struct {
		Collection string `json:"collection"`
	}

	ProxyEndpoint struct {
		Host string `json:"host"`
	}

	PageEndpoint struct {
		File string `json:"file"`
		// TODO postprocessors for markdown?
	}
)

const (
	StaticEndpointKind     = EKind("static")     // {static.path}/ -> /{collection.path}
	CollectionEndpointKind = EKind("collection") // /{collection.path} <- crud
	ProxyEndpointKind      = EKind("proxy")      // /{collection.path} -> {proxy.host}
	PageEndpointKind       = EKind("page")       // /{collection.path} <- {page.file}
)
