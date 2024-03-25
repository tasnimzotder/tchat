package cmd

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/_client/internal/display"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
	"github.com/tasnimzotder/tchat/_client/pkg/client"
	"github.com/tasnimzotder/tchat/_client/pkg/cryptography"
	"github.com/tasnimzotder/tchat/_client/pkg/models"
)

func connectionCmd(storageClient *storage.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "conn",
		Short: "conn",
		Run: func(cmd *cobra.Command, args []string) {
			connectionCmdHandler(storageClient, cmd, args)
		},
	}
}

func connectionCmdHandler(storageClient *storage.Storage, cmd *cobra.Command, args []string) {
	connectionTypes := []string{"begin", "create", "list", "remove"}

	commandPrompt := promptui.Select{
		Label: "Select a connection type",
		Items: connectionTypes,
	}

	_, connectionType, err := commandPrompt.Run()
	if err != nil {
		return
	}

	switch connectionType {
	case "begin":
		beginConnectionHandler(storageClient)
	case "create":
		createConnectionHandler(storageClient)
	case "list":
		listConnectionHandler(storageClient)
	case "remove":
		// todo: implement
	}
}

func beginConnectionHandler(storageClient *storage.Storage) {
	display.PrintMessage("info", "Beginning connection")
	display.PrintMessage("info", "This setting will store your public RSA key on the server for other users to download \nand use to send you encrypted messages.\n")

	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	display.PrintMessage("error", "Failed to create storage")
	//	return
	//}
	//
	//defer sqlite.Close()

	publicKey, err := storageClient.GetPublicRSAKey()
	if err != nil {
		display.PrintMessage("error", "Failed to get public RSA key")
		return
	}

	publicKeyBase64 := cryptography.EncodeBase64(publicKey)

	// user
	user, err := storageClient.GetLastUser()
	if err != nil {
		display.PrintMessage("error", "Failed to get user")
		return
	}

	passKey := generatePassKey(6)
	expirationTime := "3600"
	Persistence := "temporary"

	connectionRequest := client.ConnectionRequest{
		Name:        user.Name,
		UserID:      user.ID,
		PublicKey:   publicKeyBase64,
		PassKey:     passKey,
		Expiration:  expirationTime,
		Persistence: Persistence,
	}

	// start connection
	err = storageClient.API.CreateConnection(connectionRequest)
	if err != nil {
		display.PrintMessage("error", "Failed to create connection")
		return
	}

	// display
	fmt.Printf("\n")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"User ID",
		"Passkey",
		"Expiration",
		"Persistence",
	})
	table.SetBorder(false)
	table.Append([]string{
		connectionRequest.UserID,
		connectionRequest.PassKey,
		connectionRequest.Expiration,
		connectionRequest.Persistence,
	})

	table.Render()
}

func createConnectionHandler(storageClient *storage.Storage) {
	userIdPrompt := promptui.Prompt{
		Label: "Enter the user ID",
		Validate: func(input string) error {
			if len(input) < 3 {
				return fmt.Errorf("error: user ID must be at least 3 characters")
			}
			return nil
		},
	}

	userId, err := userIdPrompt.Run()
	if err != nil {
		display.PrintMessage("error", "Invalid user ID")
		return
	}

	passKeyPrompt := promptui.Prompt{
		Label: "Enter the passkey",
		Validate: func(input string) error {
			if len(input) < 6 {
				return fmt.Errorf("error: passkey must be at least 6 characters")
			}
			return nil
		},
	}

	passKey, err := passKeyPrompt.Run()
	if err != nil {
		display.PrintMessage("error", "Invalid passkey")
		return
	}

	request := client.GetConnectionRequest{
		UserID:  userId,
		PassKey: passKey,
	}

	// get connection
	connection, err := storageClient.API.GetConnection(request)
	if err != nil {
		display.PrintMessage("error", "Failed to get connection")
		return
	}

	// log.Println(connection)
	// display.PrintMessage("info", "Name: "+connection.Name)

	nicknamePrompt := promptui.Prompt{
		Label: "Enter a nickname",
		Validate: func(input string) error {
			if len(input) < 3 {
				return fmt.Errorf("error: nickname must be at least 3 characters")
			}
			return nil
		},

		Default: connection.Name,
	}

	nickname, err := nicknamePrompt.Run()
	if err != nil {
		display.PrintMessage("error", "Invalid nickname")
		return
	}

	var contact = models.Contact{
		ID:        connection.UserID,
		Name:      nickname,
		PublicKey: connection.PublicKey,
	}

	// add contact
	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	display.PrintMessage("error", "Failed to create storage")
	//	return
	//}
	//
	//defer sqlite.Close()

	err = storageClient.SaveContact(contact)
	if err != nil {
		display.PrintMessage("error", "Failed to save contact")
		return
	}

	display.PrintMessage("info", "Contact added successfully")
}

func listConnectionHandler(storageClient *storage.Storage) {
	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	display.PrintMessage("error", "Failed to create storage")
	//	return
	//}
	//
	//defer sqlite.Close()

	contacts, err := storageClient.GetContacts()
	if err != nil {
		display.PrintMessage("error", "Failed to get contacts")
		return
	}

	// display
	fmt.Printf("\n")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name",
		"ID",
	})
	table.SetBorder(false)

	for _, contact := range contacts {
		table.Append([]string{
			contact.Name,
			contact.ID,
		})
	}

	table.Render()

	// display.PrintMessage("info", "Contacts listed successfully")
}

// utils
func generatePassKey(length int) string {
	const digits = "0123456789"
	var passKey bytes.Buffer

	for i := 0; i < length; i++ {
		passKey.WriteByte(digits[rand.Intn(len(digits))])
	}

	return passKey.String()
}
