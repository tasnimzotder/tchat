package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/internal"
	"github.com/tasnimzotder/tchat/client/pkg/crypto"
	"github.com/tasnimzotder/tchat/client/pkg/file"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to the chat server",
	Run:   SendCmd,
}

func init() {
	sendCmd.Flags().StringP("recipient", "r", "", "Recipient's user name")
	sendCmd.Flags().StringP("message", "m", "", "Message to send")
	sendCmd.Flags().StringP("file", "f", "", "File to send")
}

func SendCmd(cmd *cobra.Command, args []string) {
	var err error
	var recipientName, message string
	messageType := "text"

	recipientFlagSet := cmd.Flags().Changed("recipient") // The key change
	messageFlagSet := cmd.Flags().Changed("message")
	fileFlagSet := cmd.Flags().Changed("file")

	if recipientFlagSet && (messageFlagSet || fileFlagSet) {
		recipientName, _ = cmd.Flags().GetString("recipient")

		if messageFlagSet {
			message, _ = cmd.Flags().GetString("message")
		} else if fileFlagSet {
			fileValue, _ := cmd.Flags().GetString("file")
			dir, _ := os.Getwd()

			filePath := dir + "/" + fileValue

			messageByte, err := file.GetFileContents(filePath)
			if err != nil {
				log.Printf("Failed to read file: %v", err)
				return
			}

			fileExt := filepath.Ext(filePath)

			// remove the initial dot if present
			if len(fileExt) > 0 && fileExt[0] == '.' {
				fileExt = fileExt[1:]
			}

			messageType = fmt.Sprintf("file:%s", fileExt)
			message = string(messageByte)
		}
	} else if !(recipientFlagSet || messageFlagSet || fileFlagSet) {
		// recipient name prompt
		recipientNamePrompt := promptui.Prompt{
			Label: "Recipient Name",
		}

		recipientName, err = recipientNamePrompt.Run()
		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}

		// message prompt
		messagePrompt := promptui.Prompt{
			Label: "Message",
			Validate: func(input string) error {
				if len(input) < 1 {
					return nil
				}

				return nil
			},
		}

		message, err = messagePrompt.Run()
		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}

	} else {
		cmd.Help()
		return
	}

	// get contact details
	contact, err := file.GetContactByName(recipientName)
	if err != nil {
		log.Printf("Failed to get contact: %v", err)
		return
	}

	var encryption crypto.Encryptioner = &crypto.RSAEncryption{}

	publicKey, err := file.GetPublicKeyByUserID(contact.ID)
	if err != nil {
		log.Printf("Failed to get public key: %v", err)
		return
	}

	encryptedMessage, err := encryption.EncryptMessage([]byte(message), publicKey)
	if err != nil {
		log.Printf("Failed to encrypt message: %v", err)
		return
	}

	encodedMessage := encryption.EncodeBase64(encryptedMessage)

	err = internal.SendMessage(contact.ID, messageType, encodedMessage)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return
	}

	log.Printf("Message sent successfully")
}
