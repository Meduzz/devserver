package main

import (
	"github.com/Meduzz/commando"
	"github.com/Meduzz/commando/model"
)

/*
 1. Initiate the app from the config.
 2. Create a webserver with [run] endpoints.
*/

func init() {
	run := commando.Command("run", func(ec model.ExecuteCommand) error {
		return nil
	})

	run.Description = "Run a handy app"
}
