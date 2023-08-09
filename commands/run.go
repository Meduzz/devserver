package commands

import "github.com/spf13/cobra"

func init() {
	run := &cobra.Command{}
	run.Use = "run"
	run.Example = "run {--file=app.json|--dir=apps/} [--port=8765]"
	run.Short = "start handy with either an app or a directory of apps"
	run.Flags().String("file", "app.json", "set which app (config) to load")
	run.Flags().String("dir", ".", "load all apps in a directory")
	run.Flags().Int("port", 8765, "set port that handy starts on")
	run.RunE = runHandler

	run.MarkFlagFilename("file", "json")
	run.MarkFlagDirname("dir")

	Root.AddCommand(run)
}

func runHandler(cmd *cobra.Command, args []string) error {

	/*
		TODO
		Build a app context per app we're loading
		Initiate the app server with all the app contexts
		For each app:
		- Setup databases
		- Setup endpoints
	*/
	return nil
}
