package commands

import "github.com/spf13/cobra"

var Root = &cobra.Command{
	DisableAutoGenTag:  true,
	DisableSuggestions: true,
}
