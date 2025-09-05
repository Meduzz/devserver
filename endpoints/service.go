package endpoints

import (
	"encoding/json"
	"fmt"

	"github.com/Meduzz/devserver/collections"
	"github.com/Meduzz/devserver/filedb"
	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/valyala/fasthttp"
)

/*
TODO
1. Start all services and let them register their endpoint.
2. Call endpoint.Start() to register all endpoints with webserver
3. Start webserver

We need to handle these kinds of endpoints.
- Dir (static)
  > Delegate to the webframework.
- Dynamic (collections)
  > Handle with specific code
  >> The api collection will be proxy though
- Proxy
  > Delegate to the webframework
- File
  > Delegate to the webframework
*/

type (
	endpointService struct {
		endpoints   []*model.Endpoint
		collections collections.CollectionsService
	}
)

var (
	_ model.Service    = (*endpointService)(nil)
	_ model.Controller = (*endpointService)(nil)
)

/*
TODO
*/

func NewEndpointService(endpoints []*model.Endpoint, collections collections.CollectionsService) model.Service {
	return &endpointService{
		endpoints:   endpoints,
		collections: collections,
	}
}

func (e *endpointService) Start() error {
	return nil
}

func (e *endpointService) Stop() error {
	return nil
}

func (e *endpointService) Setup(srv *fiber.App) error {
	slice.Fold(e.endpoints, nil, func(ep *model.Endpoint, agg error) error {
		if agg != nil {
			return agg
		}

		switch ep.Kind {
		case model.PageEndpointKind:
			config, err := fromJson[model.PageEndpoint](ep.Config)

			if err != nil {
				return err
			}

			srv.Get(ep.Path, func(ctx fiber.Ctx) error {
				return ctx.SendFile(config.File)
			})
		case model.ProxyEndpointKind:
			config, err := fromJson[model.ProxyEndpoint](ep.Config)

			if err != nil {
				return err
			}

			srv.All(ep.Path, proxy.Forward(config.Host, &fasthttp.Client{
				NoDefaultUserAgentHeader: true,
				DisablePathNormalizing:   true,
			}))
		case model.CollectionEndpointKind:
			config, err := fromJson[model.CollectionEndpoint](ep.Config)

			if err != nil {
				return err
			}

			col := e.collections.Lookup(config.Collection)

			switch col.Kind {
			case model.ApiCollectionKind:
				colCfg, err := fromJson[model.ApiCollection](col.Config)

				if err != nil {
					return err
				}

				srv.All(ep.Path, proxy.Forward(colCfg.Host, &fasthttp.Client{
					NoDefaultUserAgentHeader: true,
					DisablePathNormalizing:   true,
				}))
			case model.FileCollectionKind:
				colCfg, err := fromJson[model.FileCollection](col.Config)

				if err != nil {
					return err
				}

				db, err := filedb.NewFileDB(colCfg.File, colCfg.ID, col.Fields)

				if err != nil {
					return err
				}

				url := fmt.Sprintf("%s/:id", ep.Path)
				srv.Get(ep.Path, func(ctx fiber.Ctx) error {
					req := make(map[string]any)
					err := ctx.Bind().JSON(req)

					if err != nil {
						return err
					}

					rows, err := db.List()

					if err != nil {
						return err
					}

					return ctx.JSON(rows, "application/json")
				})
				srv.Post(ep.Path, func(ctx fiber.Ctx) error {
					req := make(map[string]any)
					err := ctx.Bind().JSON(req)

					if err != nil {
						return err
					}

					res, err := db.Create(req)

					if err != nil {
						return err
					}

					return ctx.JSON(res)
				})
				srv.Put(url, func(ctx fiber.Ctx) error {
					req := make(map[string]any)
					err := ctx.Bind().JSON(req)

					if err != nil {
						return err
					}

					id := ctx.Params("id", "0")
					res, err := db.Update(id, req)

					if err != nil {
						return err
					}

					return ctx.JSON(res, "application/json")
				})
				srv.Delete(url, func(ctx fiber.Ctx) error {
					id := ctx.Params("id", "0")
					return db.Delete(id)
				})
			case model.StaticCollectionKind:
				colCfg, err := fromJson[model.StaticCollection](col.Config)

				if err != nil {
					return err
				}

				srv.Get(ep.Path)
			}
		}

		return nil
	})

	return nil
}

func fromJson[T any](data json.RawMessage) (*T, error) {
	it := new(T)
	err := json.Unmarshal(data, it)

	if err != nil {
		return nil, err
	}

	return it, nil
}
