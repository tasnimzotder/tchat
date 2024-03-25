package main

import (
	"log"
	"os"

	"github.com/tasnimzotder/tchat/_client/internal/storage"

	"github.com/tasnimzotder/tchat/_client/cmd"
	"github.com/tasnimzotder/tchat/_client/pkg/client"
)

func main() {
	reqScheme := "http"
	host := "api.tchat.tasnim.dev"
	version := "v0.0.3-beta"

	apiClient := client.NewClient(
		host,
		reqScheme,
	)

	storageClient, err := storage.NewStorage(apiClient)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// migrate db
	if err := storageClient.Migrate(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// cmd.Execute(apiClient)
	if err := cmd.Execute(storageClient, version); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
