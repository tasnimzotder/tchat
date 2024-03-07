package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/internal"
	"github.com/tasnimzotder/tchat/client/pkg/file"
	"github.com/tasnimzotder/tchat/client/pkg/message"
)

var messageCmd = &cobra.Command{
	Use:   "msg",
	Short: "Message related commands",
	Long:  `Message related commands`,
	Run:   MessageCmd,
}

func init() {
	messageCmd.Flags().StringP("display", "d", "", "Display message; serial number")
	messageCmd.Flags().StringP("clear", "c", "", "Clear messages; all or last")
	messageCmd.Flags().StringP("save", "s", "", "Save message to a file")
}

func MessageCmd(cmd *cobra.Command, args []string) {
	clearFlagSet := cmd.Flags().Changed("clear") // The key change
	if clearFlagSet {
		clearValue, _ := cmd.Flags().GetString("clear")

		if clearValue == "all" {
			err := file.ClearMessagesFile()
			if err != nil {
				log.Printf("Failed to clear messages file: %v", err)
				return
			}
		}

		return
	}

	messages, err := internal.GetMessages()
	if err != nil {
		log.Printf("No new messages!")
		return
	}

	for _, message := range messages {
		err = file.AppendToMessagesFile(message)
		if err != nil {
			log.Printf("Failed to write to messages file: %v", err)
			return
		}
	}

	displayFlagSet := cmd.Flags().Changed("display") // The key change
	saveFlagSet := cmd.Flags().Changed("save")       // The key change

	if displayFlagSet {
		displayValue, _ := cmd.Flags().GetString("display")

		serialNumber, _ := strconv.Atoi(displayValue)
		// message.DisplayMessages(serialNumber)
		_message := message.DisplaySingleMessageRaw(serialNumber)

		// copy to clipboard
		// util.CopyToClipboard(_message.Payload)

		if saveFlagSet {
			fileName, _ := cmd.Flags().GetString("save")
			dir, _ := os.Getwd()

			filePath := dir + "/" + fileName
			payloadBytes := []byte(_message.Payload)

			err := file.SaveFile(filePath, payloadBytes)
			if err != nil {
				log.Printf("Failed to save file: %v", err)
				return
			}

			log.Printf("File saved to: %v", filePath)
		} else {
			fmt.Printf("%v", _message.Payload)
		}
	} else {
		message.DisplayMessages(10)
	}

}
