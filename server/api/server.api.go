package api

import (
	"encoding/json"
	"fmt"
	"github.com/tasnimzotder/tchat/server/models"
	"log"
	"net/http"
)

type ServerAPI struct {
	Server        http.Server
	MessageStacks map[string][]models.Message
}

func NewServerAPI() *ServerAPI {
	return &ServerAPI{
		Server:        http.Server{},
		MessageStacks: make(map[string][]models.Message),
	}
}

func (s *ServerAPI) Start(port string) {
	log.Println("Starting server")

	// routes
	http.HandleFunc("/ping", s.PingHandler)
	http.HandleFunc("/v1/user/create", s.createUserHandler)
	http.HandleFunc("/v1/message/send", s.sendMessageHandler)
	http.HandleFunc("/v1/message/get", s.getMessageHandler)

	// start server
	s.Server.Addr = fmt.Sprintf(":%s", port)
	err := s.Server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func (s *ServerAPI) PingHandler(w http.ResponseWriter, _ *http.Request) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: "pong",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
