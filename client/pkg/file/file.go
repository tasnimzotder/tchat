package file

import (
	"log"
	"os"
)

func getConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tchat"
}

func GetFileContents(fileName string) ([]byte, error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil, err
	}

	return contents, nil
}

func SaveFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)

	if err != nil {
		log.Printf("Failed to save file: %v", err)
		return err
	}

	return nil
}
