package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/pkg/config"
)

// StartConnection establishes a connection and sends a connection request to the server.
//
// It takes a connection request model as a parameter and returns an error.
func StartConnection(connectionRequest models.Connection) error {
	u := url.URL{
		Scheme: "http",
		Host:   config.GetEnvVariable("SERVER_HOST"),
		// Host: "localhost:5050",
		Path: "/v1/connection/set",
	}

	bodyBytes, err := json.Marshal(connectionRequest)
	if err != nil {
		return err
	}

	post, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return err
	}

	defer post.Body.Close()

	var response struct {
		Message string `json:"status"`
	}

	err = json.NewDecoder(post.Body).Decode(&response)
	if err != nil {
		return err
	}

	log.Println(response.Message)

	return nil
}

type GetConnectionRequest struct {
	UserID  string `json:"user_id"`
	PassKey string `json:"pass_key"`
}

// GetConnection retrieves a connection using the provided GetConnectionRequest.
//
// getConnectionRequest GetConnectionRequest
// Connection, error
func GetConnection(getConnectionRequest GetConnectionRequest) (models.Connection, error) {
	u := url.URL{
		Scheme: "http",
		Host:   config.GetEnvVariable("SERVER_HOST"),
		// Host: "localhost:5050",
		Path: "/v1/connection/get",
	}

	bodyBytes, err := json.Marshal(getConnectionRequest)
	if err != nil {
		return models.Connection{}, err
	}

	post, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return models.Connection{}, err
	}

	defer post.Body.Close()

	var response models.Connection

	err = json.NewDecoder(post.Body).Decode(&response)
	if err != nil {
		return models.Connection{}, err
	}

	return response, nil
}
