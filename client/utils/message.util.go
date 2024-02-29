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

	for i := len(messages) - 1; i >= len(messages)-limit; i-- {
		displayMessage(table, messages[i])
	}

	table.Render()

}

func displayMessage(table *tablewriter.Table, message models.Message) {
	//fmt.Println("Sender ID \t Timestamp | Message")
	//fmt.Println("--------- | --------- | -------")
	//fmt.Printf("%s \t %s | %s\n", message.SenderID, message.Timestamp, message.Payload)

	// trim message if it's too long
	if len(message.Payload) > 25 {
		message.Payload = message.Payload[:25] + "..."
	}

	// sender id, show only last 5 characters
	if len(message.SenderID) > 5 {
		message.SenderID = "..." + message.SenderID[len(message.SenderID)-5:]

	}

	table.Append([]string{message.SenderID, message.Timestamp, message.Payload})
}