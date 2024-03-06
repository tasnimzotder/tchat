package message

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/pkg/crypto"
	"github.com/tasnimzotder/tchat/client/pkg/file"
)

func DisplayMessages(limit int) {
	messages, err := file.ReadFromMessagesFile()
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
		_plainTextMessage := plainTextMessage(messages[i].Payload)
		if _plainTextMessage != "" {
			messages[i].Payload = _plainTextMessage
		}

		// if i == len(messages)-1 {
		// 	util.CopyToClipboard(messages[i].Payload)
		// }

		displayMessage(table, messages[i])
	}

	// insert a new line
	fmt.Println()
	table.Render()
}

func plainTextMessage(payload string) string {
	var encryption crypto.Encryptioner = &crypto.RSAEncryption{}

	privateKey, err := file.GetPrivateKey()
	if err != nil {
		fmt.Println("Failed to get private key")
		return ""
	}

	decodedBytes, err := encryption.DecodeBase64(payload)
	if err != nil {
		fmt.Println("Failed to decode message")
		return ""
	}

	decryptedBytes, err := encryption.DecryptMessage(decodedBytes, privateKey)
	if err != nil {
		fmt.Println("Failed to decrypt message")
		return ""
	}

	return string(decryptedBytes)
}

func displayMessage(table *tablewriter.Table, message models.Message) {
	//var encryption crypto.Encryptioner = &crypto.RSAEncryption{}

	// get contact details
	contact, err := file.GetContactByID(message.SenderID)
	if err != nil {
		fmt.Println("Failed to get contact details")
		return
	}

	// trim message if it's too long
	if len(message.Payload) > 15 {
		message.Payload = message.Payload[:15] + "..."
	}

	// contact name, show only first 6 characters
	if len(contact.Name) > 6 {
		contact.Name = contact.Name[:6] + "..."

	}

	table.Append([]string{contact.Name, message.Timestamp, message.MessageType, message.Payload})
}

func DisplaySingleMessageRaw(serialNumber int) models.Message {
	messages, err := file.ReadFromMessagesFile()
	if err != nil {
		fmt.Println("No messages to display")
		return models.Message{}
	}

	message := messages[len(messages)-serialNumber]

	_plainTextMessage := plainTextMessage(message.Payload)

	if _plainTextMessage != "" {
		message.Payload = _plainTextMessage
	}

	return message
}
