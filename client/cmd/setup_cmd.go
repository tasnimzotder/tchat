package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/_client/internal/display"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
	"github.com/tasnimzotder/tchat/_client/pkg/cryptography"
)

func setupCmd(storageClient *storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "setup",
		Short: "setup",
		Run: func(cmd *cobra.Command, args []string) {
			setupCmdHandler(storageClient, cmd, args)
		},
	}
}

func setupCmdHandler(storageClient *storage.Storage, cmd *cobra.Command, args []string) {
	display.PrintMessage("info", "Setting up tchat")
	fmt.Printf("\n")

	// todo: check if user already exists

	namePrompt := promptui.Prompt{
		Label: "Enter your name",
		Validate: func(input string) error {
			if len(input) < 3 {
				return fmt.Errorf("error: name must be at least 3 characters")
			}
			return nil
		},
	}

	name, err := namePrompt.Run()
	if err != nil {
		display.PrintMessage("error", "Invalid name")
		return
	}

	passwordPrompt := promptui.Prompt{
		Label: "Enter your password",
		Mask:  '*',
		Validate: func(input string) error {
			if len(input) < 6 {
				return fmt.Errorf("error: password must be at least 6 characters")
			}
			return nil
		},
	}

	password, err := passwordPrompt.Run()
	if err != nil {
		display.PrintMessage("error", "Invalid password")
		return
	}

	createNewUser(storageClient, name, password)
}

func createNewUser(storageClient *storage.Storage, name, password string) {
	user, err := storageClient.API.CreateUser(name, password)
	if err != nil {
		display.PrintMessage("error", "Failed to create user")
		return
	}

	// if err := file.StoreRSAKeys(privateKey, publicKey); err != nil {
	// 	display.PrintMessage("error", "Failed to store RSA keys")
	// 	return
	// }

	// save user details to the db
	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	display.PrintMessage("error", "Failed to create database")
	//	return
	//}
	//
	//defer sqlite.Close()

	// crypto
	privateKey, publicKey, err := cryptography.GenerateKeyPair(2048)
	if err != nil {
		display.PrintMessage("error", "Failed to generate key pair")
		return
	}

	// convert privateKey and publicKey to []byte
	privateKeyBytes, publicKeyBytes := cryptography.ConvertRSAToBytes(privateKey, publicKey)

	// delete all previous keys
	if err = storageClient.DeleteRSAKeys(); err != nil {
		display.PrintMessage("error", "Failed to delete RSA keys")
		return
	}

	// save keys to db as []byte
	if err = storageClient.SaveRSAKeys(privateKeyBytes, publicKeyBytes); err != nil {
		display.PrintMessage("error", "Failed to save RSA keys")
		return
	}

	// create all previous users in the db
	users, _ := storageClient.GetAllUsers()
	for _, u := range users {
		// delete all users
		if err := storageClient.DeleteUser(u); err != nil {
			display.PrintMessage("error", "Failed to delete user")
			return
		}
	}

	err = storageClient.SaveUser(user)
	if err != nil {
		display.PrintMessage("error", "Failed to save user")
		return
	}

	display.PrintMessage("info", "User created successfully")
}
