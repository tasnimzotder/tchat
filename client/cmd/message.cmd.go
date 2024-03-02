package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/client/services"
	"github.com/tasnimzotder/tchat/client/utils"
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
			err := utils.ClearMessagesFile()
			if err != nil {
				log.Printf("Failed to clear messages file: %v", err)
				return
			}
		}

		return
	}

	messages, err := services.GetMessages()
	if err != nil {
		log.Printf("No new messages!")
		return
	}

	for _, message := range messages {
		err = utils.AppendToMessagesFile(message)
		if err != nil {
			log.Printf("Failed to write to messages file: %v", err)
			return
		}
	}

	displayFlagSet := cmd.Flags().Changed("display") // The key change
	if displayFlagSet {
		displayValue, _ := cmd.Flags().GetString("display")

		limit, _ := strconv.Atoi(displayValue)
		utils.DisplayMessages(limit)
	} else {
		utils.DisplayMessages(10)
	}

}
