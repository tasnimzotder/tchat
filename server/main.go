package main

import "github.com/tasnimzotder/tchat/server/api"

func main() {
	server := api.NewServerAPI()
	server.Start()
}
