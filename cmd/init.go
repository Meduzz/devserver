package cmd

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/registry"
	"github.com/Meduzz/devserver/cms"
	handy "github.com/Meduzz/devserver/model"
	"github.com/Meduzz/devserver/services"
	"github.com/spf13/cobra"
)

/*
 1. Craete an empty app with the provided name.
 2. Start a webserver with the new app as base with [cms] endpoints.
 3. Redirect to CMS.
*/

func init() {
	i := commando.Command("init", func(c *cobra.Command, s []string) error {
		name, err := c.Flags().GetString("name")

		if err != nil {
			return err
		}

		cms := cms.NewCMS("handy.json", &handy.App{
			Name: name,
		})
		services.Register(cms)

		port, err := c.Flags().GetInt("port")

		if err != nil {
			return err
		}

		// TODO open browser "async"

		// blocking
		return services.Start(port)
	})

	i.Description = "Initialize a handy app"
	i.AddFlag(flags.IntFlag("port", 8080, "Start the handy server on this port"))
	i.AddFlag(flags.StringFlag("name", "", "The name of the app that gets initiated"))

	registry.RegisterCommand(i)
}
