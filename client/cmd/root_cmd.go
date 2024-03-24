package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/_client/pkg/client"
)



func Execute(apiClient *client.Client) error {
	var rootCmd = &cobra.Command{
		Use:   "tchat",
		Short: "..",
	}

	rootCmd.AddCommand(setupCmd(apiClient))
	rootCmd.AddCommand(connectionCmd(apiClient))
	rootCmd.AddCommand(messageCmd(apiClient))

	return rootCmd.Execute()
}
