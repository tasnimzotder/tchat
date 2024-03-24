package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tasnimzotder/tchat/server/middleware"
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
	router := http.NewServeMux()

	// routes
	router.HandleFunc("POST /v1/user/create", s.createUserHandler)
	router.HandleFunc("POST /v1/message/send", s.sendMessageHandler)
	router.HandleFunc("GET /v1/message/get", s.getMessageHandler)

	// connection
	router.HandleFunc("POST /v1/connection/set", s.setConnectionHandler)
	router.HandleFunc("POST /v1/connection/get", s.getConnectionHandler)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	router.HandleFunc("GET /health", s.healthCheckHandler)
	router.HandleFunc("GET /ping", s.PingHandler)

	// middleware
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + port,
		Handler: stack(router),
	}

	log.Printf("Server started on port %s", port)
	if err := server.ListenAndServe(); err != nil {
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

	health := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	_ = json.NewEncoder(w).Encode(health)
}
