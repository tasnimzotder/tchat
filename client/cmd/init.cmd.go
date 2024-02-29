package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/services"
)

var userName string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new chat server",
	Run:   InitCmd,
}

func InitCmd(cmd *cobra.Command, args []string) {
	services.InitializeNewConnection(userName)
}

func init() {
	initCmd.Flags().StringVarP(&userName, "username", "u", "", "Your username")
}
