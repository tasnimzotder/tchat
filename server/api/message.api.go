package api

import (
	"encoding/json"
	"github.com/tasnimzotder/tchat/server/models"
	"github.com/tasnimzotder/tchat/server/utils"
	"net/http"
	"time"
)

func (s *ServerAPI) sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ErrorResponse(w, "method not allowed", nil)

		return
	}

	var message models.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorResponse(w, "", err)

		return
	}

	if message.SenderID == "" || message.RecipientID == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorResponse(w, "sender or receiver id is missing", nil)

		return
	}

	defer r.Body.Close()

	message.Timestamp = time.Now().Format(time.RFC3339)

	// store message in the message stack
	s.MessageStacks[message.RecipientID] = append(s.MessageStacks[message.RecipientID], message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(struct {
		Response string         `json:"response"`
		Message  models.Message `json:"message"`
	}{
		Response: "message sent successfully",
		Message:  message,
	})
}

func (s *ServerAPI) getMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ErrorResponse(w, "method not allowed", nil)

		return
	}

	var recipientID = r.URL.Query().Get("id")

	if recipientID == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorResponse(w, "recipient id is missing", nil)

		return
	}

	var messages []models.Message

	for _, message := range s.MessageStacks[recipientID] {
		messages = append(messages, message)
	}

	// clear the message stack
	s.MessageStacks[recipientID] = []models.Message{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct {
		Response string           `json:"response"`
		Messages []models.Message `json:"messages"`
	}{
		Response: "messages fetched successfully",
		Messages: messages,
	})
}
