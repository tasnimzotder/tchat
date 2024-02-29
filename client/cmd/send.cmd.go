package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/services"
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
}

func SendCmd(cmd *cobra.Command, args []string) {
	if recipientID == "" {
		cmd.Help()
		return
	}

	messageType := "text"

	err := services.SendMessage(recipientID, message, messageType)
	if err != nil {
		return
	}
}
