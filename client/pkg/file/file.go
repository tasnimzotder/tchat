package file

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	DatabasePath string
	Name         string
	ID           string
}

func GetProjectStoragePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return homeDir + "/.config/tchat"
}

func CreateConfigFile(filePath string, config *Config) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("database_path = " + config.DatabasePath + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("name = " + config.Name + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("id = " + config.ID + "\n")
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid config line: " + line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "database_path":
			config.DatabasePath = value
		case "name":
			config.Name = value
		case "id":
			config.ID = value
		default:
			return nil, errors.New("unknown config key: " + key)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func SaveFile(filePath string, data []byte) error {
	path := GetProjectStoragePath() + "/" + filePath

	// create the directory if it doesn't exist
	if err := os.MkdirAll(GetProjectStoragePath(), 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetRootFilesInDir(dirPath string) ([]string, error) {
	fileInfos, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		files = append(files, fileInfo.Name())
	}

	return files, nil
}

func FileInfo(filePath string) fs.FileInfo {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Failed to get file size: %v", err)
		return nil
	}

	return fileInfo
}

func Extension(filePath string) string {
	fileExt := filepath.Ext(filePath)

	// remove the initial dot if present
	if len(fileExt) > 0 && fileExt[0] == '.' {
		fileExt = fileExt[1:]
	}

	return fileExt
}
