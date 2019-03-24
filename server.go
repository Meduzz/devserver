package main

import (
	"os"

	"./plugins"
	"./server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "devserver"
	app.Version = "180316"
	app.Description = "A multitool web server"

	app.Commands = []cli.Command{
		server.NewServer(plugins.Pages()),
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
