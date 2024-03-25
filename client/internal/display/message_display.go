package display

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
	"github.com/tasnimzotder/tchat/_client/pkg/cryptography"
	"github.com/tasnimzotder/tchat/_client/pkg/models"
	"github.com/tasnimzotder/tchat/_client/pkg/utils"
)

func DisplayMessages(storageClient *storage.Storage, messages []models.Message) {
	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}
	//
	//defer sqlite.Close()

	privateKey, err := storageClient.GetPrivateRSAKey()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var decodedMessages []models.Message

	for _, msg := range messages {
		_decodedMsg := decodeMessage(msg, privateKey)

		// update sender name (from id)
		contact, err := storageClient.GetContactByID(_decodedMsg.SenderID)
		if err == nil {
			_decodedMsg.SenderID = contact.Name
		}

		decodedMessages = append(decodedMessages, _decodedMsg)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Sender",
		"Timestamp",
		"Type",
		"Mode",
		"Size",
		"Payload",
	})
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	//table.SetColWidth(1000)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
	)

	for _, msg := range decodedMessages {
		appendMessage(table, msg)
	}

	table.Render()
}

func decodeMessage(msg models.Message, key []byte) models.Message {
	rsaKey, err := cryptography.ConvertPrivateBytesToRSA(key)
	if err != nil {
		return msg
	}

	decodedPayload, err := cryptography.DecodeBase64(msg.Payload)
	if err != nil {
		return msg
	}

	decryptedPayload, err := cryptography.DecryptMessage(decodedPayload, rsaKey)
	if err != nil {
		return msg
	}

	msg.Payload = string(decryptedPayload)
	return msg
}

func appendMessage(table *tablewriter.Table, message models.Message) {
	// trim message if it's too long
	if len(message.Payload) > 15 {
		message.Payload = message.Payload[:15] + "..."
	}

	// contact name, show only first 6 characters
	// if len(contact.Name) > 6 {
	// 	contact.Name = contact.Name[:6] + "..."

	// }

	if len(message.SenderID) > 6 {
		message.SenderID = message.SenderID[:6] + "..."
	}

	fileSize := utils.BytesToSize(message.FileSize)

	table.Append([]string{
		message.SenderID,
		message.Timestamp,
		message.MessageType,
		message.FileMode,
		fileSize,
		message.Payload,
	})
}

func DisplaySingleMessage(storageClient *storage.Storage, message models.Message) {
	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}
	//
	//defer sqlite.Close()

	privateKey, err := storageClient.GetPrivateRSAKey()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	decodedMessage := decodeMessage(message, privateKey)

	fmt.Printf("\n")
	fmt.Println(decodedMessage.Payload)
	fmt.Printf("\n")
}
