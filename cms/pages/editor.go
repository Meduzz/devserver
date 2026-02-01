package pages

import (
	"encoding/json"
	"fmt"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/logic"
	"github.com/Meduzz/gml/tags"
)

type (
	optionBacker struct {
		value    string
		label    string
		selected bool
	}
)

func Editor(app *model.App, ep *model.Endpoint) gml.Tag {
	title := fmt.Sprintf("[%s] CMS Editor", app.Name)
	static := &model.StaticEndpoint{}
	proxy := &model.ProxyEndpoint{}

	if ep == nil {
		ep = &model.Endpoint{
			Kind: model.StaticEndpointKind,
		}
	} else {
		switch ep.Kind {
		case model.ProxyEndpointKind:
			json.Unmarshal(ep.Config, proxy)
		case model.StaticEndpointKind:
			json.Unmarshal(ep.Config, static)
		}
	}

	kinds := make([]*optionBacker, 0)
	kinds = append(kinds,
		backer("", "- Select -", false),
		backer(string(model.ProxyEndpointKind), "Proxy", ep.Kind == model.ProxyEndpointKind),
		backer(string(model.StaticEndpointKind), "Static", ep.Kind == model.StaticEndpointKind))

	form := tags.Form(gml.Tags(
		// name
		text("Name:", "name", ep.Name, "My endpoint", "true"),
		// kind
		zelect("Kind:", "kind", "kind", kinds),
		// path
		text("Path:", "path", ep.Path, "/api", "true"),
		// drop prefix
		text("Drop prefix:", "drop", ep.Drop, "/prefix", "true"),
		// config
		// - host
		text("Host:", "host", proxy.Host, "http://host:port", "kind == 'proxy'"),
		// - dir
		text("Dir:", "dir", static.Dir, "public/", "kind == 'static'"),
		// - spa
		text("SPA:", "spa", static.SPA, "index.html", "kind == 'static'"),
		tags.Div(
			tags.Button(gml.Text("Save"), attr.Type("submit")),
		),
	), gml.StringAttribute("method", "POST"), gml.StringAttribute("x-data", fmt.Sprintf("{kind: '%s'}", ep.Kind)))

	return page(title, form)
}

func text(label, name, value, placeholder, show string) gml.Tag {
	return tags.Div(gml.Tags(
		tags.Label(gml.Text(label)),
		tags.Input(gml.Empty(), attr.Type("text"), attr.Name(name), attr.Value(value), attr.Placeholder(placeholder)),
	), gml.StringAttribute("x-show", show))
}

func zelect(label, name, bind string, options []*optionBacker) gml.Tag {
	return tags.Div(gml.Tags(
		tags.Label(gml.Text(label)),
		tags.Select(
			logic.Slice(options, option, tags.Option(gml.Text("- Empty - "))),
			attr.Name(name),
			gml.StringAttribute("x-model", bind),
		),
	))
}

func option(it *optionBacker) gml.Tag {
	return tags.Option(gml.Text(it.label), attr.Value(it.value), gml.AnyAttribute("selected", it.selected))
}

func backer(value, label string, selected bool) *optionBacker {
	return &optionBacker{
		value:    value,
		label:    label,
		selected: selected,
	}
}
