package endpoints

import (
	"os"
	"path"
	"strings"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/gin-gonic/gin"
)

func (e *EndpointService) serveStatic() gin.HandlerFunc {
	staticEndpoints := slice.Filter(e.endpoints, func(it *endpoint) bool {
		return it.kind == model.StaticEndpointKind
	})

	return func(ctx *gin.Context) {
		fileName := ctx.Request.URL.Path

		// find the best match
		bestMatch := slice.Fold(staticEndpoints, nil, func(in, agg *endpoint) *endpoint {
			if !strings.HasPrefix(fileName, in.path) {
				return agg
			}

			if agg == nil {
				return in
			}

			if len(in.path) > len(agg.path) {
				return in
			}

			return agg
		})

		// if there's no match, then 404
		if bestMatch == nil {
			// 404
			ctx.AbortWithStatus(404)
			return
		} else {
			// else apply middleware(s) (drop prefix)
			if bestMatch.drop != "" {
				fileName = strings.TrimPrefix(fileName, bestMatch.drop)
			}
		}

		root := bestMatch.static.dir
		file, err := os.Open(path.Join(root, fileName))

		if err != nil {
			if os.IsNotExist(err) {
				if bestMatch.static.index != "" {
					ctx.File(path.Join(root, bestMatch.static.index))
					return
				}

				ctx.AbortWithStatus(404)
				return
			}

			ctx.AbortWithStatus(500)
			return
		}

		file.Close()

		ctx.File(path.Join(root, fileName))
	}
}
