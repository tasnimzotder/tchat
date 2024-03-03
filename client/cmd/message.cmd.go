package cmd

import (
	"log"
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
	messageCmd.Flags().StringP("display", "d", "", "Display messages")
	messageCmd.Flags().StringP("clear", "c", "", "Clear messages; all or last")
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
	if displayFlagSet {
		displayValue, _ := cmd.Flags().GetString("display")

		limit, _ := strconv.Atoi(displayValue)
		message.DisplayMessages(limit)
	} else {
		message.DisplayMessages(10)
	}

}
