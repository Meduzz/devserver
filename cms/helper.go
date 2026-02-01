package cms

import (
	"github.com/Meduzz/gml"
	"github.com/gin-gonic/gin"
)

func render(ctx *gin.Context, it gml.Tag) {
	ctx.Header("Content-Type", "text/html")
	ctx.String(200, it.Render())
}
