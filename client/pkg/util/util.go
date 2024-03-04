package util

import (
	"log"

	"github.com/atotto/clipboard"
)

func CopyToClipboard(text string) {
	err := clipboard.WriteAll(text)
	if err != nil {
		log.Fatalf("Failed to copy to clipboard: %v", err)
	}
}
