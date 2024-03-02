package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tasnimzotder/tchat/client/models"
)

func DisplayMessages(limit int) {
	messages, err := ReadFromMessagesFile()
	if err != nil {
		//log.Printf("Failed to read from messages file: %v", err)
		fmt.Println("No messages to display")
		return
	}

	// reverse messages
	// todo: reverse messages

	// display messages
	if limit > len(messages) {
		limit = len(messages)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sender", "Timestamp", "Type", "Message"})
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	//table.SetColWidth(1000)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
	)

	for i := len(messages) - 1; i >= len(messages)-limit; i-- {
		displayMessage(table, messages[i])
	}

	// insert a new line
	fmt.Println()
	table.Render()
}

func displayMessage(table *tablewriter.Table, message models.Message) {
	// get contact details
	contact, err := GetContactByID(message.SenderID)
	if err != nil {
		fmt.Println("Failed to get contact details")
		return
	}

	privateKey, err := GetPrivateKey()
	if err != nil {
		fmt.Println("Failed to get private key")
		return
	}

	decodedBytes, err := DecodeBase64(message.Payload)
	if err != nil {
		fmt.Println("Failed to decode message")
		return
	}

	decryptedBytes, err := DecryptMessage(decodedBytes, privateKey)
	if err != nil {
		fmt.Println("Failed to decrypt message")
		return
	}

	message.Payload = string(decryptedBytes)

	// trim message if it's too long
	if len(message.Payload) > 15 {
		message.Payload = message.Payload[:32] + "..."
	}

	// contact name, show only first 6 characters
	if len(contact.Name) > 6 {
		contact.Name = contact.Name[:6] + "..."

	}

	table.Append([]string{contact.Name, message.Timestamp, message.MessageType, message.Payload})
}
