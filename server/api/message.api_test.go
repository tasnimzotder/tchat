package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/tchat/server/models"
)

func TestSendMessageHandler(t *testing.T) {
	// Create a new instance of the ServerAPI
	s := &ServerAPI{
		MessageStacks: make(map[string][]models.Message),
	}

	// Create a test message
	message := models.Message{
		SenderID:    "senderID",
		RecipientID: "recipientID",
		MessageType: "text",
		Payload:     "Hello, world!",
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	// Encode the message as JSON
	messageJSON, _ := json.Marshal(message)

	// Create a new HTTP request with the encoded message as the request body
	req, err := http.NewRequest(http.MethodPost, "/v1/message/send", bytes.NewBuffer(messageJSON))
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Call the sendMessageHandler function
	s.sendMessageHandler(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, res.Code)

	// Decode the response body
	var response struct {
		Response string         `json:"response"`
		Message  models.Message `json:"message"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the response fields
	assert.Equal(t, "message sent successfully", response.Response)
	assert.Equal(t, message.SenderID, response.Message.SenderID)
	assert.Equal(t, message.RecipientID, response.Message.RecipientID)
	assert.Equal(t, message.Payload, response.Message.Payload)
	// assert.WithinDuration(t, time.Now(), response.Message.Timestamp, time.Second)
}

func TestGetMessageHandler(t *testing.T) {
	// Create a new instance of the ServerAPI
	s := &ServerAPI{
		MessageStacks: make(map[string][]models.Message),
	}

	// Add test messages to the message stack
	recipientID := "recipientID"
	message1 := models.Message{
		SenderID:    "senderID1",
		RecipientID: recipientID,
		MessageType: "text",
		Payload:     "Hello, world!",
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	message2 := models.Message{
		SenderID:    "senderID2",
		RecipientID: recipientID,
		MessageType: "text",
		Payload:     "How are you?",
		Timestamp:   time.Now().Format(time.RFC3339),
	}
	s.MessageStacks[recipientID] = []models.Message{message1, message2}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/v1/message?id="+recipientID, nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Call the getMessageHandler function
	s.getMessageHandler(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Decode the response body
	var messages []models.Message
	err = json.NewDecoder(res.Body).Decode(&messages)
	assert.NoError(t, err)

	// Check the response fields
	assert.Equal(t, 2, len(messages))
	assert.Equal(t, message1.SenderID, messages[0].SenderID)
	assert.Equal(t, message1.RecipientID, messages[0].RecipientID)
	assert.Equal(t, message1.Payload, messages[0].Payload)
	assert.Equal(t, message2.SenderID, messages[1].SenderID)
	assert.Equal(t, message2.RecipientID, messages[1].RecipientID)
	assert.Equal(t, message2.Payload, messages[1].Payload)

	// Check if the message stack is cleared
	assert.Equal(t, 0, len(s.MessageStacks[recipientID]))
}
