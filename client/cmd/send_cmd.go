package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/_client/internal/display"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
	"github.com/tasnimzotder/tchat/_client/pkg/client"
	"github.com/tasnimzotder/tchat/_client/pkg/cryptography"
	"github.com/tasnimzotder/tchat/_client/pkg/file"
	"github.com/tasnimzotder/tchat/_client/pkg/models"
	"github.com/tasnimzotder/tchat/_client/pkg/utils"
)

func sendMessageCmd(apiClient *client.Client) *cobra.Command {
	sendMessageCmd := &cobra.Command{
		Use:   "send",
		Short: "Send message",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sendMessageHandler(apiClient, cmd, args)
		},
	}

	return sendMessageCmd
}

func sendMessageHandler(apiClient *client.Client, _ *cobra.Command, args []string) error {
	var err error
	var recipientName, messageType, messageText string
	var currDir, fileName string

	if len(args) == 0 {
		recipientName, err = promptForRecipientName()
		if err != nil {
			return err
		}

		messageType, err = promptForMessageType()
		if err != nil {
			return err
		}

		if messageType == "text" {
			messageTextPrompt := promptui.Prompt{
				Label: "Enter message text",
				Validate: func(input string) error {
					if len(input) == 0 {
						return fmt.Errorf("message text cannot be empty")
					}
					return nil
				},
			}

			messageText, err = messageTextPrompt.Run()
			if err != nil {
				return err
			}
		} else if messageType == "file" {
			currDir, fileName, err = promptForFileName()
			if err != nil {
				return err
			}
		}

	} else {
		// todo: implement
	}

	// contact details
	sqlite, err := storage.NewSQLiteStorage()
	if err != nil {
		return err
	}

	defer sqlite.Close()

	contact, err := sqlite.GetContactByName(recipientName)
	if err != nil {
		return err
	}

	// get sender id
	senderID, err := sqlite.GetUserID()
	if err != nil {
		return err
	}

	var messageData []byte

	if messageType == "file" {
		filePath := currDir + "/" + fileName
		messageData, err = file.ReadFile(filePath)
		if err != nil {
			return err
		}
	} else if messageType == "text" {
		messageData = []byte(messageText)
	}

	// encrypt message
	encryptedMessage, err := cryptography.EncryptMessageWithRSAString(messageData, contact.PublicKey)
	if err != nil {
		return err
	}

	fileInfo := file.FileInfo(currDir + "/" + fileName)

	// check file size > 5MB
	if messageType == "file" && fileInfo.Size() > 1*1024*1024 {
		display.PrintMessage("error", "File size cannot exceed 1MB")
		return nil
	}

	if messageType == "file" {
		messageType = fmt.Sprintf("file/%s", file.Extension(fileName))
	}

	msgReq := models.Message{
		MessageType: messageType,
		Payload:     encryptedMessage,
		SenderID:    senderID,
		RecipientID: contact.ID,
		Timestamp:   utils.GetCurrTimeStr(),
		FileExt:     file.Extension(fileName),
		FileSize:    fileInfo.Size(),
		FileName:    fileInfo.Name(),
		FileMode:    fileInfo.Mode().String(),
	}

	if messageType == "text" {
		msgReq.FileName = ""
		msgReq.FileSize = 0
		msgReq.FileMode = ""
		msgReq.FileExt = ""
	}

	// fmt.Println("Message request:", msgReq)

	table := tablewriter.NewWriter(os.Stdout)

	if messageType == "text" {
		table.SetHeader([]string{
			"Recipient",
			"Message Type",
			"Message",
		})

		table.Append([]string{
			recipientName,
			msgReq.MessageType,
			trimText(messageText, 23),
		})
	} else {
		table.SetHeader([]string{
			"Recipient",
			"Message Type",
			"File Name",
			"File Size",
			"File Mode",
		})

		table.Append([]string{
			recipientName,
			msgReq.MessageType,
			msgReq.FileName,
			utils.BytesToSize(msgReq.FileSize),
			msgReq.FileMode,
		})
	}

	table.SetBorder(false)
	table.Render()
	fmt.Printf("\n")

	// send message
	err = apiClient.SendMessage(msgReq)

	if err != nil {
		display.PrintMessage("error", "Failed to send message")
		return err
	}

	display.PrintMessage("info", "Message sent successfully")
	return nil
}

func trimText(text string, length int) string {
	if len(text) > length {
		return text[:length] + "..."
	}

	return text
}

// prompts
func promptForRecipientName() (string, error) {
	sqlite, err := storage.NewSQLiteStorage()
	if err != nil {
		return "", err
	}

	defer sqlite.Close()

	contacts, err := sqlite.GetContacts()
	if err != nil {
		return "", err
	}

	var recipientNames []string
	for _, contact := range contacts {
		recipientNames = append(recipientNames, contact.Name)
	}

	recipientPrompt := promptui.Select{
		Label: "Select a recipient",
		Items: recipientNames,
	}

	_, recipientName, err := recipientPrompt.Run()
	if err != nil {
		return "", err
	}

	// fmt.Println("Selected recipient:", recipientName)

	return recipientName, nil
}

func promptForMessageType() (string, error) {
	types := []string{"text", "file"}
	messageTypePrompt := promptui.Select{
		Label: "Select message type",
		Items: types,
	}

	_, messageType, err := messageTypePrompt.Run()
	if err != nil {
		return "", err
	}

	return messageType, nil
}

func promptForFileName() (currDir, fileName string, err error) {
	// get current directory
	currDir, err = os.Getwd()
	if err != nil {
		return "", "", err
	}

	// get files in current directory
	files, err := file.GetRootFilesInDir(currDir)
	if err != nil {
		return "", "", err
	}

	// check if there are any files
	if len(files) == 0 {
		return "", "", fmt.Errorf("no files found in %s", currDir)
	}

	filePrompt := promptui.Select{
		Label: "Select a file",
		Items: files,
	}

	_, fileName, err = filePrompt.Run()
	if err != nil {
		return "", "", err
	}

	return currDir, fileName, nil
}
