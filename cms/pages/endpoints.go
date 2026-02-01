package pages

import (
	"fmt"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/htmx"
	"github.com/Meduzz/gml/logic"
	"github.com/Meduzz/gml/tags"
)

func Endpoints(app *model.App) gml.Tag {
	title := fmt.Sprintf("[%s] CMS Endpoints", app.Name)
	add := tags.Div(
		tags.A(gml.Text("+"), attr.Href("/editor")),
	)
	list := logic.Slice(app.Endpoints, endpointListItem, tags.P(gml.Text("No endpoints.")))

	return page(title, gml.Tags(add, list))
}

func endpointListItem(ep *model.Endpoint) gml.Tag {
	return tags.Div(
		tags.A(gml.Tags(
			tags.Span(gml.Text(ep.Name)),
			tags.Span(gml.Text(ep.Path)),
			tags.Span(gml.Text(ep.Drop)),
			tags.Span(gml.Text(fmt.Sprintf("[%s]", ep.Kind))),
			tags.Button(gml.Text("Remove"), attr.Type("button"), htmx.Confirm("Really remove this endpoint?"), htmx.Delete(fmt.Sprintf("/endpoint/%s", ep.Name)), htmx.Target("body")),
		), attr.Href(fmt.Sprintf("/editor/%s", ep.Name))),
	)
}
