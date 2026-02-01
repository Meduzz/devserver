package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/registry"
	"github.com/Meduzz/devserver/app"
	"github.com/Meduzz/devserver/endpoints"
	"github.com/Meduzz/devserver/model"
	"github.com/Meduzz/devserver/services"
	"github.com/spf13/cobra"
)

func init() {
	run := commando.Command("run", func(c *cobra.Command, s []string) error {
		port, err := c.Flags().GetInt("port")

		if err != nil {
			return err
		}

		config, err := c.Flags().GetString("config")

		if err != nil {
			return err
		}

		// load app
		app, err := loadApp(config)

		if err != nil {
			return err
		}

		svc := endpoints.NewEndpointService()

		// figure out endpoints
		svc.Offer("", app)

		// iterate child apps
		for key, config := range app.Children {
			app, err = loadApp(config)

			if err != nil {
				return err
			}

			svc.Offer(key, app)
		}

		services.Register(svc)

		return services.Start(port)
	})

	run.Description = "Run a handy app"
	run.AddFlag(flags.IntFlag("port", 8080, "Start the handy server on this port"))
	run.AddFlag(flags.StringFlag("config", "handy.json", "Load the handy config from this URI"))

	registry.RegisterCommand(run)
}

func loadApp(configFile string) (*model.App, error) {
	return app.Load(configFile)
}
