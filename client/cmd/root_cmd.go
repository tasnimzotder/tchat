package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
)

func Execute(storageClient *storage.Storage, versionStr string) error {
	var rootCmd = &cobra.Command{
		Use:   "tchat",
		Short: "tchat is a terminal messaging app",
		Run: func(cmd *cobra.Command, args []string) {
			//	parse flag
			version, _ := cmd.Flags().GetBool("version")
			if version {
				_ = printVersion(versionStr)
			}

			help, _ := cmd.Flags().GetBool("help")
			if help {
				_ = cmd.Help()
			}

			//	if no flag is passed, print help
			if !version && !help {
				_ = cmd.Help()
			}
		},
	}

	//check if user setup
	exists, err := storageClient.IsUserExist()
	if err != nil {
		return err
	}

	if !exists {
		setupCmdHandler(storageClient, nil, nil)
	}

	rootCmd.AddCommand(setupCmd(storageClient))
	rootCmd.AddCommand(connectionCmd(storageClient))
	rootCmd.AddCommand(messageCmd(storageClient))

	// flag for version
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print version")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")

	return rootCmd.Execute()
}

func printVersion(version string) error {
	fmt.Println(version)
	return nil
}
