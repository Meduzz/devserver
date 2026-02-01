package pages

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/tags"
)

func page(title string, body gml.Tag) gml.Tag {
	return tags.Html(gml.Tags(
		tags.Head(gml.Tags(
			tags.Title(gml.Text(title)),
			tags.Script(gml.Text(""), attr.Src("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"), attr.Defer()),
			tags.Script(gml.Text(""), attr.Src("https://cdn.jsdelivr.net/npm/htmx.org@2.0.8/dist/htmx.min.js"), gml.StringAttribute("crossorigin", "anonymous")),
		)),
		tags.Body(gml.Tags(
			tags.H1(gml.Text(title)),
			body,
		)),
	))
}
