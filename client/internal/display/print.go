package display

import "fmt"

func PrintMessage(msg_type, msg string) {
	prefix := ""
	color := ""

	switch msg_type {
	case "info":
		prefix = "[INFO]"
		color = "\033[32m" // green
	case "error":
		prefix = "[ERROR]"
		color = "\033[31m" // red
	case "warning":
		prefix = "[WARNING]"
		color = "\033[33m" // yellow
	}

	// square block
	msg_icon := "\u25A0"

	fmt.Printf("%s %s %s %s\n", color+msg_icon, prefix, msg, "\033[0m")
}
