package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tasnimzotder/tchat/server/models"
)

type ServerAPI struct {
	Server           http.Server
	MessageStacks    map[string][]models.Message
	ConnectionStacks map[string]models.Connection
}

func NewServerAPI() *ServerAPI {
	return &ServerAPI{
		Server:           http.Server{},
		MessageStacks:    make(map[string][]models.Message),
		ConnectionStacks: make(map[string]models.Connection),
	}
}

func (s *ServerAPI) Start(port string) {
	log.Println("Starting server")

	// routes
	http.HandleFunc("/health", s.healthCheckHandler)
	http.HandleFunc("/ping", s.PingHandler)
	http.HandleFunc("/v1/user/create", s.createUserHandler)
	http.HandleFunc("/v1/message/send", s.sendMessageHandler)
	http.HandleFunc("/v1/message/get", s.getMessageHandler)

	// connection
	http.HandleFunc("/v1/connection/set", s.setConnectionHandler)
	http.HandleFunc("/v1/connection/get", s.getConnectionHandler)

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

func (s *ServerAPI) healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	log.Printf("Health check request received")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}
