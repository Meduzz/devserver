package endpoints

import (
	"net/http"
	"net/url"

	"github.com/Meduzz/devserver/config"
	"github.com/Meduzz/devserver/lib/common"
	"github.com/Meduzz/devserver/lib/handlers/model"
	"github.com/gin-gonic/gin"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
)

type (
	proxy struct {
		ctx *common.AppContext
		cfg *config.Proxy
	}
)

func NewProxy(ctx *common.AppContext, cfg *config.Proxy) model.Endpoint {
	return &proxy{ctx, cfg}
}

func (p *proxy) Handler() (gin.HandlerFunc, error) {
	if p.cfg.Kind == "http" {
		return httpProxy(p.cfg.Http)
	} else {
		return gin.WrapF(http.NotFound), nil
	}
}

func (p *proxy) Context() *common.AppContext {
	return p.ctx
}

func httpProxy(cfg *config.HttpProxy) (gin.HandlerFunc, error) {
	it, err := forward.New(forward.PassHostHeader(true))

	if err != nil {
		return nil, err
	}

	rr, err := roundrobin.New(it)

	if err != nil {
		return nil, err
	}

	theUrl, err := url.Parse(cfg.Target)

	if err != nil {
		return nil, err
	}

	rr.UpsertServer(theUrl)

	return gin.WrapH(rr), nil
}
