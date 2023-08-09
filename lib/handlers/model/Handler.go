package model

import (
	"github.com/Meduzz/devserver/config"
	"github.com/Meduzz/devserver/lib/common"
	"github.com/gin-gonic/gin"
)

type (
	Datasource interface {
		Context() *common.AppContext
		Handler() (gin.HandlerFunc, error)
	}

	Endpoint interface {
		Context() *common.AppContext
		Handler() (gin.HandlerFunc, error)
	}

	DatasourceFactory func(*config.Datasource) Datasource
	EndpointFactory   func(*config.Endpoint) Endpoint
)
