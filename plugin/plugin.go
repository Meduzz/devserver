package plugin

import (
	"github.com/gin-gonic/gin"
)

type (
	Plugin interface {
		Prefix() string
		Name() string
		Bind(*gin.RouterGroup)
	}
)
