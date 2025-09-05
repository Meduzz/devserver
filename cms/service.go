package cms

import (
	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/gofiber/fiber/v3"
)

/*
TODO
*/

var (
	_ model.Service    = (*cms)(nil)
	_ model.Controller = (*cms)(nil)
)

func NewCMS(app *model.App) model.Service {
	return &cms{
		app: app,
	}
}

func (c *cms) Start() error {
	println("CMS starting")
	return nil
}

func (c *cms) Stop() error {
	println("CMS stopping")
	return nil
}

func (c *cms) Setup(srv *fiber.App) error {
	// TODO static

	// setup endpoints
	epApi := srv.Group("/api/endpoints")
	coApi := srv.Group("/api/collections")

	epApi.Get("", c.listEndpoints)
	epApi.Post("", c.createEndpoint)
	epApi.Put("", c.updateEndpoint)
	epApi.Delete("/:name", c.removeEndponit)

	coApi.Get("", c.listCollections)
	coApi.Post("", c.createCollection)
	coApi.Put("", c.updateCollection)
	coApi.Delete("/:name", c.removeCollection)

	return nil
}

func (c *cms) listEndpoints(ctx fiber.Ctx) error {
	return ctx.JSON(c.app.Endpoints)
}

func (c *cms) createEndpoint(ctx fiber.Ctx) error {
	req := &model.Endpoint{}
	err := ctx.Bind().JSON(req)

	if err != nil {
		return err
	}

	c.app.Endpoints = append(c.app.Endpoints, req)

	return ctx.JSON(req)
}

func (c *cms) updateEndpoint(ctx fiber.Ctx) error {
	req := &model.Endpoint{}
	err := ctx.Bind().JSON(req)

	if err != nil {
		return err
	}

	c.app.Endpoints = slice.Map(c.app.Endpoints, func(ep *model.Endpoint) *model.Endpoint {
		if ep.Name == req.Name {
			return req
		} else {
			return ep
		}
	})

	return ctx.JSON(req)
}

func (c *cms) removeEndponit(ctx fiber.Ctx) error {
	name := ctx.Params("name")
	c.app.Endpoints = slice.Filter(c.app.Endpoints, func(e *model.Endpoint) bool {
		return e.Name != name
	})

	return nil
}

func (c *cms) listCollections(ctx fiber.Ctx) error {
	return ctx.JSON(c.app.Collections)
}

func (c *cms) createCollection(ctx fiber.Ctx) error {
	req := &model.Collection{}
	err := ctx.Bind().JSON(req)

	if err != nil {
		return err
	}

	c.app.Collections = append(c.app.Collections, req)

	return ctx.JSON(req)
}

func (c *cms) updateCollection(ctx fiber.Ctx) error {
	req := &model.Collection{}
	err := ctx.Bind().JSON(req)

	if err != nil {
		return err
	}

	c.app.Collections = slice.Map(c.app.Collections, func(ep *model.Collection) *model.Collection {
		if ep.Name == req.Name {
			return req
		} else {
			return ep
		}
	})

	return ctx.JSON(req)
}

func (c *cms) removeCollection(ctx fiber.Ctx) error {
	name := ctx.Params("name")
	c.app.Collections = slice.Filter(c.app.Collections, func(e *model.Collection) bool {
		return e.Name != name
	})

	return nil
}
