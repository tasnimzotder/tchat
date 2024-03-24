package file

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
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
