package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tasnimzotder/tchat/_client/internal/display"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
)

func messageCmd(storageClient *storage.Storage) *cobra.Command {
	_messageCmd := &cobra.Command{
		Use:   "msg",
		Short: "Message commands",
		Run: func(cmd *cobra.Command, args []string) {
			messageCmdHandler(storageClient, cmd, args)
		},
	}

	_messageCmd.Flags().BoolP("clear", "c", false, "Clear messages")
	_messageCmd.Flags().BoolP("display", "d", false, "Display nth Message (default: last)")
	_messageCmd.Flags().BoolP("save", "s", false, "Save nth Message (default: last)")

	_messageCmd.AddCommand(sendMessageCmd(storageClient))
	// messageCmd.AddCommand(getMessagesCmd(apiClient))

	return _messageCmd
}

func messageCmdHandler(storageClient *storage.Storage, cmd *cobra.Command, args []string) error {
	clearFlag, err := cmd.Flags().GetBool("clear")
	if err != nil {
		return err
	}

	displayFlag, err := cmd.Flags().GetBool("display")
	if err != nil {
		return err
	}

	saveFlag, err := cmd.Flags().GetBool("save")
	if err != nil {
		return err
	}

	if saveFlag {
		// todo: implement
		fmt.Println("Saving last message")
	}

	//sqlite, err := storage.NewSQLiteStorage()
	//if err != nil {
	//	return err
	//}
	//
	//defer sqlite.Close()

	userID, err := storageClient.GetUserID()
	if err != nil {
		return err
	}

	messages, err := storageClient.API.GetMessages(userID)
	if err != nil {
		return err
	}

	if clearFlag {
		err = storageClient.DeleteMessages()

		if err != nil {
			return err
		}

		return nil
	}

	// save messages
	err = storageClient.SaveMessages(messages)
	if err != nil {
		return err
	}

	// get messages
	messages, err = storageClient.GetMessages()
	if err != nil {
		return err
	}

	// reverse messages
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	if displayFlag {
		// todo: implement
		idx := 1

		if len(args) > 0 {
			_idx := args[0]

			// convert string to int
			idx, err = strconv.Atoi(_idx)
			if err != nil {
				return err
			}

			if idx < 1 {
				return fmt.Errorf("index cannot be less than 1")
			}

		}

		if idx > len(messages) {
			display.PrintMessage("error", "Index out of range")

			return fmt.Errorf("index out of range")
		}

		message := messages[len(messages)-idx]
		display.DisplaySingleMessage(storageClient, message)
	} else {
		display.DisplayMessages(storageClient, messages)
	}

	return nil
}
