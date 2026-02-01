package services

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Meduzz/helper/service"
	"github.com/Meduzz/helper/service/web"
	"github.com/gin-gonic/gin"
)

func init() {
	web.SetEngine(server)
}

var (
	server *gin.Engine = gin.Default()
)

func Register(it service.Service) {
	service.AddService(it)
}

func Start(port int) error {
	// create the webserver
	err := service.Start()

	if err != nil {
		return err
	}

	go handleShutdown()

	return server.Run(fmt.Sprintf(":%d", port))
}

func Stop() error {
	return service.Stop()
}

func handleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	<-c
	err := Stop()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
