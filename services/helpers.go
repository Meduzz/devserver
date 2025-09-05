package services

import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/helper/fp/slice"
	"github.com/gofiber/fiber/v3"
)

/*
TODO
* Switch from gin to echo
*/

var (
	services []model.Service
	server   *fiber.App
)

func Register(it model.Service) {
	services = append(services, it)
}

func Start(port int) error {
	// create the webserver
	server = fiber.New()

	go handleShutdown()

	// iterate the registered services and start them
	err := slice.Fold(services, nil, func(in model.Service, agg error) error {
		if agg != nil {
			return agg
		}

		err := in.Start()

		if err != nil {
			return err
		}

		controller, ok := in.(model.Controller)

		if ok {
			controller.Setup(server)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return server.Listen(fmt.Sprintf(":%d", port))
}

func Stop() error {
	return slice.Fold(services, nil, func(in model.Service, agg error) error {
		err := in.Stop()

		if err != nil {
			if agg != nil {
				return errors.Join(agg, err)
			}

			return err
		}

		return agg
	})
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
