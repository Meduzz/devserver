package commands

import "github.com/spf13/cobra"

var Root = &cobra.Command{}

func init() {
	Root.Use = "handy"
	Root.CompletionOptions.DisableDefaultCmd = true
}
