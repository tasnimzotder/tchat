package cmd

import (
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/pkg/file"
)

var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Contact related commands",
	Long:  `Contact related commands`,
	Run:   ContactCmd,
}

func ContactCmd(cmd *cobra.Command, args []string) {
	// contact subcommands
	// name prompt
	namePrompt := promptui.Prompt{
		Label: "Contact Name",
		Validate: func(input string) error {
			if len(input) < 3 {
				return nil
			}
			return nil
		},
	}

	name, err := namePrompt.Run()
	if err != nil {
		log.Printf("Failed to get contact name: %v", err)
		return
	}

	//	id prompt
	idPrompt := promptui.Prompt{
		Label: "Contact ID",
		Validate: func(input string) error {
			if len(input) < 36 {
				return nil
			}
			return nil
		},
	}

	id, err := idPrompt.Run()
	if err != nil {
		log.Printf("Failed to get contact id: %v", err)
		return
	}

	//	key prompt
	keyPrompt := promptui.Prompt{
		Label: "RSA Public Key File Name",
		//	todo: validate the key
	}

	key, err := keyPrompt.Run()
	if err != nil {
		log.Printf("Failed to get contact key: %v", err)
		return
	}

	// process the contact
	dir, _ := os.Getwd()
	keyPath := dir + "/" + key

	destKeyPath, err := file.StoreContactPublicKey(id, keyPath)
	if err != nil {
		log.Printf("Failed to store contact public key: %v", err)
		return
	}

	contact := models.Contact{
		ID:   id,
		Name: name,
		Key:  destKeyPath,
	}

	err = file.WriteToContactFile(contact)
	if err != nil {
		log.Printf("Failed to write to contact file: %v", err)
		return
	}

	log.Printf("user: %s, name: %s, key: %s", id, name, key)

	log.Printf("Contact added successfully")
}
