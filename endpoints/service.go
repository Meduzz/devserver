package endpoints

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/result"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/service"
	"github.com/Meduzz/helper/service/web"
	"github.com/gin-gonic/gin"
)

type (
	endpoint struct {
		kind   model.EndpointKind
		path   string
		drop   string
		proxy  *proxy
		static *static
	}

	proxy struct {
		host url.URL
	}

	static struct {
		dir   string
		index string
	}

	EndpointService struct {
		endpoints []*endpoint
	}
)

var (
	_ service.Service = &EndpointService{}
	_ web.WebApi      = &EndpointService{}
)

func NewEndpointService() *EndpointService {
	return &EndpointService{}
}

func (e *EndpointService) Start() error {
	return nil
}

func (e *EndpointService) Stop() error {
	return nil
}

func (e *EndpointService) Setup(srv *gin.Engine) {
	staticRegistered := false
	slice.ForEach(e.endpoints, func(ep *endpoint) {
		switch ep.kind {
		case model.ProxyEndpointKind:
			srv.Any(ep.path, proxyTo(ep))
		case model.StaticEndpointKind:
			if !staticRegistered {
				srv.NoRoute(e.serveStatic())
				staticRegistered = true
			}
		}
	})
}

func (e *EndpointService) Offer(appDir string, app *model.App) error {
	ops := result.Batch(app.Endpoints, func(ep *model.Endpoint) (*endpoint, error) {
		switch ep.Kind {
		case model.StaticEndpointKind:
			s, err := loadStatic(appDir, ep)

			if err != nil {
				return nil, err
			}

			path := ep.Path

			return &endpoint{
				kind:   ep.Kind,
				path:   path,
				static: s,
				drop:   ep.Drop,
			}, nil
		case model.ProxyEndpointKind:
			p, err := loadProxy(ep)

			if err != nil {
				return nil, err
			}

			path := ep.Path

			if path == "" || path == "/" {
				path = "/*file"
			}

			if path != "/" && !strings.HasSuffix(path, "/*file") {
				path = fmt.Sprintf("%s/*file", path)
			}

			return &endpoint{
				kind:  ep.Kind,
				path:  path,
				proxy: p,
				drop:  ep.Drop,
			}, nil
		default:
			return nil, fmt.Errorf("unknown endpoint kind: %s", ep.Kind)
		}
	})

	eps, err := ops.Get()

	if err != nil {
		return err
	}

	e.endpoints = append(e.endpoints, eps...)

	return nil
}

func loadStatic(appDir string, ep *model.Endpoint) (*static, error) {
	cfg := &model.StaticEndpoint{}
	err := json.Unmarshal(ep.Config, cfg)

	if err != nil {
		return nil, err
	}

	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	it := &static{
		dir:   path.Join(wd, appDir, cfg.Dir),
		index: cfg.SPA,
	}

	return it, nil
}

func loadProxy(ep *model.Endpoint) (*proxy, error) {
	cfg := &model.ProxyEndpoint{}
	err := json.Unmarshal(ep.Config, cfg)

	if err != nil {
		return nil, err
	}

	downstream, err := url.Parse(cfg.Host)

	if err != nil {
		return nil, err
	}

	it := &proxy{
		host: *downstream,
	}

	return it, nil
}
