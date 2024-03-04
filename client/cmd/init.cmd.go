package cmd

import (
	"errors"
	"log"

	"github.com/tasnimzotder/tchat/client/internal"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new chat server",
	Run:   InitCmd,
}

func InitCmd(cmd *cobra.Command, args []string) {
	// init subcommands
	// name prompt
	promptName := promptui.Prompt{
		Label: "What is your name?",
		Validate: func(s string) error {
			if len(s) < 3 {
				return errors.New("name must be at least 3 characters long")
			}

			return nil
		},
	}

	name, err := promptName.Run()
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return
	}

	internal.InitializeNewConnection(name)
}
