package main

import (
	"github.com/tasnimzotder/tchat/_client/internal/storage"
	"log"
	"os"

	"github.com/tasnimzotder/tchat/_client/cmd"
	"github.com/tasnimzotder/tchat/_client/pkg/client"
)

func init() {
	//err := godotenv.Load()
	//
	//if err != nil {
	//	os.Setenv("TC_SERVER_HOST", "api.tchat.tasnim.dev")
	//}
}

func main() {
	reqScheme := "http"
	host := "api.tchat.tasnim.dev"
	version := "v0.0.2"

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
