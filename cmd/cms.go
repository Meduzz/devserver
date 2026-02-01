package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/registry"
	"github.com/Meduzz/devserver/app"
	"github.com/Meduzz/devserver/cms"
	"github.com/Meduzz/devserver/services"
	"github.com/spf13/cobra"
)

/*
 1. Try to load and initiate the app from the config. Dont allow urls.
 2. Create a webserver with [cms] endpoints.
 3. Open the browser to the CMS root endpoint.
 4. On exit, save app.
*/

func init() {
	cms := commando.Command("cms", func(c *cobra.Command, s []string) error {
		configLocation, err := c.Flags().GetString("config")

		if err != nil {
			return err
		}

		app, err := app.LoadFile(configLocation)

		if err != nil {
			return err
		}

		cms := cms.NewCMS(configLocation, app)
		services.Register(cms)

		port, err := c.Flags().GetInt("port")

		if err != nil {
			return err
		}

		// TODO open browser "async"

		// blocking
		return services.Start(port)
	})

	cms.Description = "Enter CMS mode with a handy app"
	cms.AddFlag(flags.IntFlag("port", 8080, "Start the handy server on this port"))
	cms.AddFlag(flags.StringFlag("config", "handy.json", "Load the handy config from this URI"))

	registry.RegisterCommand(cms)
}
