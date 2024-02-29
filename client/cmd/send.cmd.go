package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/services"
	"github.com/tasnimzotder/tchat/client/utils"
	"log"
	"os"
)

var (
	recipientID string
	message     string
)

var sendCmd = &cobra.Command{
	Use:   "s",
	Short: "Send a message to the chat server",
	Run:   SendCmd,
}

func init() {
	sendCmd.Flags().StringVarP(&recipientID, "recipient", "r", "", "Recipient's user ID")
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "Message to send")
	sendCmd.Flags().StringP("file", "f", "", "File to send")
}

func SendCmd(cmd *cobra.Command, args []string) {
	if recipientID == "" {
		cmd.Help()
		return
	}

	//var messageByte []byte
	messageType := "text"

	fileFlagSet := cmd.Flags().Changed("file") // The key change
	if fileFlagSet {
		fileValue, _ := cmd.Flags().GetString("file")
		dir, _ := os.Getwd()
		messageType = "file"

		filePath := dir + "/" + fileValue

		messageByte, err := utils.GetFileContents(filePath)
		if err != nil {
			log.Printf("Failed to read file: %v", err)
			return
		}

		// todo: encrypt the file contents
		cypherByte, err := utils.EncryptMessage(messageByte)
		if err != nil {
			log.Printf("Failed to encrypt file: %v", err)
			return
		}

		encodedData := utils.EncodeBase64(cypherByte)

		message = encodedData
	} else {
		messageType = "text"
	}

	err := services.SendMessage(recipientID, messageType, message)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return
	}

	log.Printf("Message sent successfully")
}
