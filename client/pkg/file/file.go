package file

import (
	"log"
	"os"
)

// getConfigDir returns the configuration directory path.
//
// No parameters.
// Returns a string.
func getConfigDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tchat"
}

// GetFileContents reads the contents of a file and returns them as a byte slice.
// It takes the file name as input and returns the file contents and any error encountered.
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
