package server

import (
	"github.com/Meduzz/devserver/bootstrap"
	"github.com/gin-gonic/gin"
)

type (
	HttpServer struct {
		srv *gin.Engine
	}
)

func NewHttpServer() *HttpServer {
	srv := gin.Default()
	return &HttpServer{srv}
}

func (h *HttpServer) CreateEndpoints(endpoints []*bootstrap.EndpointDTO) {

}

func (h *HttpServer) Run(hostPort string) error {
	return h.srv.Run(hostPort)
}
