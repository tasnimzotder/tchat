package file

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tasnimzotder/tchat/client/models"
)

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
