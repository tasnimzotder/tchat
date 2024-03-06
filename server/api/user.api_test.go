package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/tchat/server/models"
)

func TestCreateUserHandler(t *testing.T) {
	// Create a new server API instance
	s := &ServerAPI{}

	// Create a request body
	user := models.User{
		ID:   "123",
		Name: "John Doe",
	}

	body, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/v1/user/create", bytes.NewBuffer(body))
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler function
	s.createUserHandler(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Check the response body
	var responseUser models.User
	err = json.Unmarshal(rr.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
}
