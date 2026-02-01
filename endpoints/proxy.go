package endpoints

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
)

func proxyTo(ep *endpoint) gin.HandlerFunc {
	rp := httputil.NewSingleHostReverseProxy(&ep.proxy.host)
	if ep.drop != "" {
		rp.Director = director(rp.Director, ep.drop)
	}

	return gin.WrapH(rp)
}

func director(original func(*http.Request), prefix string) func(*http.Request) {
	return func(r *http.Request) {
		if original != nil {
			original(r)
		}
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
	}
}
