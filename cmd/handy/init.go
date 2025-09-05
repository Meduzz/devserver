package main

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/flags"
	"github.com/Meduzz/commando/model"
	"github.com/Meduzz/devserver/cms"
	handy "github.com/Meduzz/devserver/model"
	"github.com/Meduzz/devserver/services"
)

/*
 1. Start a webserver with [init, cms] endpoints.
 2. Open the browser at the init enpoint.
 3. On init post, create an app with provided name.
 4. Redirect to CMS.
*/

/*
TODO
*/

func init() {
	i := commando.Command("init", func(ec model.ExecuteCommand) error {
		cms := cms.NewCMS(&handy.App{})
		services.Register(cms)

		port, err := ec.Int("port")

		if err != nil {
			return err
		}

		// blocking
		return services.Start(port)
	})

	i.Description = "Initialize a handy app"
	i.AddFlag(flags.IntFlag("port", 8080, "Start the handy server on this port"))
}
