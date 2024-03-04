package api

import (
	"encoding/json"
	"net/http"

	"github.com/tasnimzotder/tchat/server/models"
	"github.com/tasnimzotder/tchat/server/utils"
)

func (s *ServerAPI) setConnectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ErrorResponse(w, "method not allowed", nil)

		return
	}

	var connectionRequest models.Connection

	err := json.NewDecoder(r.Body).Decode(&connectionRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorResponse(w, "", err)

		return
	}

	// add connection to stack
	s.ConnectionStacks[connectionRequest.UserID] = connectionRequest

	//  todo: persistence logic

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct {
		Message string `json:"status"`
	}{
		Message: "connection request received",
	})
}

func (s *ServerAPI) getConnectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ErrorResponse(w, "method not allowed", nil)

		return
	}

	var connectionRequest struct {
		UserID  string `json:"user_id"`
		PassKey string `json:"pass_key"`
	}

	err := json.NewDecoder(r.Body).Decode(&connectionRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorResponse(w, "", err)

		return
	}

	connection := s.ConnectionStacks[connectionRequest.UserID]

	if connection.PassKey != connectionRequest.PassKey {
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorResponse(w, "unauthorized", nil)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(connection)
}
