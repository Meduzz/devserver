package cms

import (
	"encoding/json"
	"fmt"

	"github.com/Meduzz/devserver/app"
	"github.com/Meduzz/devserver/cms/pages"
	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/service"
	"github.com/Meduzz/helper/service/web"
	"github.com/gin-gonic/gin"
)

var (
	_ service.Service = (*cms)(nil)
	_ web.WebApi      = (*cms)(nil)
)

func NewCMS(configFile string, app *model.App) service.Service {
	return &cms{
		app:        app,
		configFile: configFile,
	}
}

func (c *cms) Start() error {
	println("CMS starting")
	return nil
}

func (c *cms) Stop() error {
	println("CMS stopping")
	return app.Store(c.app, c.configFile)
}

func (c *cms) Setup(srv *gin.Engine) {
	srv.GET("/", c.endpointsPage)               // gml
	srv.GET("/editor", c.endpointsEditor)       // gml
	srv.GET("/editor/:name", c.endpointsEditor) // gml

	srv.POST("/editor", c.endpointCreate)
	srv.POST("/editor/:name", c.endpointUpdate)

	srv.DELETE("/endpoint/:name", c.endpointRemove)
}

func (c *cms) endpointsPage(ctx *gin.Context) {
	tag := pages.Endpoints(c.app)
	render(ctx, tag)
}

func (c *cms) endpointsEditor(ctx *gin.Context) {
	name := ctx.Param("name")
	var ep *model.Endpoint

	if name != "" {
		// load ep
		ep = slice.Head(slice.Filter(c.app.Endpoints, func(e *model.Endpoint) bool {
			return e.Name == name
		}))
	}

	tag := pages.Editor(c.app, ep)
	render(ctx, tag)
}

func (c *cms) endpointCreate(ctx *gin.Context) {
	ep := &model.Endpoint{}
	proxy := &model.ProxyEndpoint{}
	static := &model.StaticEndpoint{}

	ep.Name = ctx.PostForm("name")
	ep.Kind = model.EndpointKind(ctx.PostForm("kind"))
	ep.Path = ctx.PostForm("path")
	ep.Drop = ctx.PostForm("drop")

	proxy.Host = ctx.PostForm("host")

	static.Dir = ctx.PostForm("dir")
	static.SPA = ctx.PostForm("spa")

	if ep.Kind == model.ProxyEndpointKind {
		bs, _ := json.Marshal(proxy)
		ep.Config = bs
	} else {
		bs, _ := json.Marshal(static)
		ep.Config = bs
	}

	c.app.Endpoints = append(c.app.Endpoints, ep)
	ctx.Header("HX-Redirect", "/")
	ctx.Redirect(303, "/")
}

func (c *cms) endpointUpdate(ctx *gin.Context) {
	name := ctx.Param("name")

	ep := slice.Head(slice.Filter(c.app.Endpoints, func(e *model.Endpoint) bool {
		return e.Name == name
	}))

	if ep == nil {
		ctx.AbortWithError(404, fmt.Errorf("no endpoint with name %s exists", name))
		return
	}

	ep.Path = ctx.PostForm("path")
	ep.Drop = ctx.PostForm("drop")

	if ep.Kind == model.ProxyEndpointKind {
		proxy := &model.ProxyEndpoint{}
		json.Unmarshal(ep.Config, proxy)

		proxy.Host = ctx.PostForm("host")

		bs, _ := json.Marshal(proxy)
		ep.Config = bs
	} else {
		static := &model.StaticEndpoint{}
		json.Unmarshal(ep.Config, static)

		static.Dir = ctx.PostForm("dir")
		static.SPA = ctx.PostForm("spa")

		bs, _ := json.Marshal(static)
		ep.Config = bs
	}

	ctx.Header("HX-Redirect", "/")
	ctx.Redirect(303, "/")
}

func (c *cms) endpointRemove(ctx *gin.Context) {
	name := ctx.Param("name")

	c.app.Endpoints = slice.Filter(c.app.Endpoints, func(e *model.Endpoint) bool {
		return e.Name != name
	})

	ctx.Header("HX-Redirect", "/")
	ctx.Redirect(303, "/")
}
