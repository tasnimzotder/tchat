package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tchat",
	Short: "tchat is a simple terminal chat application",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(messageCmd)
}
