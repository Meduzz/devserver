package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	server := &cobra.Command{}
	server.Use = "server --hostPort=:12345"
	server.Flags().String("hostPort", ":12345", "--hostPort=:12345")
	server.Short = "Start a local dev server"
	server.RunE = serverHandler

	Root.AddCommand(server)
}

func serverHandler(cmd *cobra.Command, args []string) error {
	return nil
}
