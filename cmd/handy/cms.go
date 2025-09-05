package main

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/devserver/app"
	"github.com/Meduzz/devserver/cms"
	"github.com/Meduzz/devserver/services"
)

/*
 1. Try to load and initiate the app from the config. Dont allow urls.
 2. Create a webserver with [cms] endpoints.
 3. Open the browser to the CMS root endpoint.
 4. On ticker, save app?
*/

/*
TODO
*/

func init() {
	cms := commando.Command("cms", func(ec model.ExecuteCommand) error {
		configLocation, err := ec.String("config")

		if err != nil {
			return err
		}

		app, err := app.Load(configLocation)

		if err != nil {
			return err
		}

		cms := cms.NewCMS(app)
		services.Register(cms)

		port, err := ec.Int("port")

		if err != nil {
			return err
		}

		// blocking
		return services.Start(port)
	})

	cms.Description = "Enter CMS mode with a handy app"
	cms.AddFlag(flags.IntFlag("port", 8080, "Start the handy server on this port"))
	cms.AddFlag(flags.StringFlag("config", "handy.yaml", "Load the handy config from this URI"))
}
