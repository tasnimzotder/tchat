package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tasnimzotder/tchat/_client/cmd"
	"github.com/tasnimzotder/tchat/_client/internal/storage"
	"github.com/tasnimzotder/tchat/_client/pkg/client"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	sqlite, err := storage.NewSQLiteStorage()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// migrate db
	if err := sqlite.Migrate(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	reqScheme := "http"

	apiClient := client.NewClient(
		os.Getenv("TC_SERVER_HOST"),
		reqScheme,
	)

	// cmd.Execute(apiClient)
	if err := cmd.Execute(apiClient); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
