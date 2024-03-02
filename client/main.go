package main

import (
	"github.com/joho/godotenv"
	"github.com/tasnimzotder/tchat/client/cmd"
	"log"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		//log.Printf("Failed to load env file: %v", err)
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
