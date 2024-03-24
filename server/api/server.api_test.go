package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/tchat/server/models"
)

func TestNewServerAPI(t *testing.T) {
	serverAPI := NewServerAPI()

	// Assert that the Server field is initialized correctly
	assert.NotNil(t, serverAPI.Server)
	assert.Equal(t, http.Server{}, serverAPI.Server)

	// Assert that the MessageStacks field is initialized correctly
	assert.NotNil(t, serverAPI.MessageStacks)
	assert.Equal(t, make(map[string][]models.Message), serverAPI.MessageStacks)

	// Assert that the ConnectionStacks field is initialized correctly
	assert.NotNil(t, serverAPI.ConnectionStacks)
	assert.Equal(t, make(map[string]models.Connection), serverAPI.ConnectionStacks)
}

func TestPingHandler(t *testing.T) {
	// Create a new request with a nil body
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	serverAPI := NewServerAPI()

	// Call the PingHandler function
	handler := http.HandlerFunc(serverAPI.PingHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	expected := `{"message":"pong"}`
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n")) // 
}

func TestHealthCheckHandler(t *testing.T) {
	// Create a new request with a nil body
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	serverAPI := NewServerAPI()

	// Call the HealthCheckHandler function
	handler := http.HandlerFunc(serverAPI.healthCheckHandler)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	expected := `{"status":"ok"}`
	assert.Equal(t, expected, strings.TrimSuffix(rr.Body.String(), "\n"))
}
