package utils

import (
	"fmt"
	"time"
)

func GetCurrTimeStr() string {
	return time.Now().Format(time.RFC3339)
}

func BytesToSize(bytes int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}

	var size string
	var i int

	for i = 0; bytes >= 1024 && i < len(sizes)-1; i++ {
		bytes /= 1024
	}

	size = fmt.Sprintf("%d%s", bytes, sizes[i])

	return size
}
