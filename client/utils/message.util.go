package utils

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tasnimzotder/tchat/client/models"
	"os"
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
	table.SetHeader([]string{"Sender ID", "Timestamp", "Message"})
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	//table.SetColWidth(1000)

	for i := len(messages) - 1; i >= len(messages)-limit; i-- {
		displayMessage(table, messages[i])
	}

	table.Render()

}

func displayMessage(table *tablewriter.Table, message models.Message) {
	decodedBytes, err := DecodeBase64(message.Payload)
	if err != nil {
		fmt.Println("Failed to decode message")
		return
	}

	decryptedBytes, err := DecryptMessage(decodedBytes)
	if err != nil {
		fmt.Println("Failed to decrypt message")
		return
	}

	message.Payload = string(decryptedBytes)

	// trim message if it's too long
	if len(message.Payload) > 15 {
		message.Payload = message.Payload[:32] + "..."
	}

	// sender id, show only last 5 characters
	if len(message.SenderID) > 5 {
		message.SenderID = "..." + message.SenderID[len(message.SenderID)-5:]
	}

	table.Append([]string{message.SenderID, message.Timestamp, message.Payload})
}
