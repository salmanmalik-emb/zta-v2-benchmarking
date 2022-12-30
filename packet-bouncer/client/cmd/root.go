package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "client",
		Short:        "",
		Long:         "",
		SilenceUsage: true,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.AddCommand(startCmd, stopCmd)
}
