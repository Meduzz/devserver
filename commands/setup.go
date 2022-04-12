package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	setup := &cobra.Command{}
	setup.Use = "setup --bootstrap=<host:port>"
	setup.Flags().String("bootstrap", "", "--bootstrap=<host:port>")
}

func setupHandler(cmd *cobra.Command, args []string) error {
	url, err := cmd.Flags().GetString("bootstrap")

	if err != nil {
		return err
	}

	if url == "" {
		return fmt.Errorf("bootstrap must be set")
	}

}
