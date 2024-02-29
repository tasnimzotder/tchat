package utils

import (
	"encoding/json"
	"github.com/tasnimzotder/tchat/client/models"
	"log"
	"os"
)

func getConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tchat"
}

func ClearMessagesFile() error {
	fileName := getConfigDir() + "/messages.json"

	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json empty array
	err = json.NewEncoder(file).Encode([]models.Message{})

	return nil
}

func ReadFromMessagesFile() ([]models.Message, error) {
	fileName := getConfigDir() + "/messages.json"

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// read from file as json
	var messages []models.Message
	err = json.NewDecoder(file).Decode(&messages)
	if err != nil {
		//log.Printf("Failed to decode messages: %v", err)
		return nil, err
	}

	return messages, nil
}

func AppendToMessagesFile(message models.Message) error {
	fileName := getConfigDir() + "/messages.json"

	// create if not exists
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// create the directory

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	messages, err := ReadFromMessagesFile()
	if err != nil {
		log.Printf("Failed to read from messages file: %v", err)
		return err
	}

	messages = append(messages, message)

	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json
	err = json.NewEncoder(file).Encode(messages)
	if err != nil {
		log.Printf("Failed to encode messages: %v", err)
		return err
	}

	return nil
}

func WriteToConfigFile(config models.Config) error {
	fileName := getConfigDir() + "/config.json"

	// create if not exists
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// create the directory

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer file.Close()
	}
	// open file
	file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	// clear the file before writing
	err = file.Truncate(0)
	if err != nil {
		log.Printf("Failed to truncate file: %v", err)
		return err
	}

	// write to file as json
	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromConfigFile() (models.Config, error) {
	fileName := getConfigDir() + "/config.json"

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		return models.Config{}, err
	}

	defer file.Close()

	// read from file as json
	var config models.Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return models.Config{}, err
	}

	return config, nil
}
