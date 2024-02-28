package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ConnStruct struct {
	Conn   *websocket.Conn
	UserID string
}

type ServerAPI struct {
	Server      http.Server
	Connections map[string]ConnStruct
}

func NewServerAPI() *ServerAPI {
	return &ServerAPI{
		Server:      http.Server{},
		Connections: make(map[string]ConnStruct),
	}
}

func (s *ServerAPI) PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write([]byte("Pong"))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	} else {
		log.Printf("Response written: %v", write)
	}
}

func (s *ServerAPI) Start() {
	log.Println("Starting server")

	// routes
	http.HandleFunc("/ping", s.PingHandler)
	http.HandleFunc("/v1/message", s.messageHandler)
	http.HandleFunc("/v1/user/create", s.createUserHandler)

	// start server
	s.Server.Addr = ":8080"
	err := s.Server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
