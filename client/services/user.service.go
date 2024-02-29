package services

import (
	"bytes"
	"encoding/json"
	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/utils"
	"io"
	"log"
	"net/http"
	"net/url"
)

func InitializeNewConnection(username string) {
	body := models.User{
		Name: username,
	}

	// convert to io.Reader
	bodyBytes, err := json.Marshal(body)

	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/v1/user/create",
	}

	post, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
			return
		}
	}(post.Body)

	// read response
	var response models.User
	err = json.NewDecoder(post.Body).Decode(&response)
	if err != nil {
		log.Printf("Failed to decode response: %v", err)
		return
	}

	// save user id to config
	err = utils.WriteToConfigFile(models.Config{
		ID:   response.ID,
		Name: response.Name,
	})

	if err != nil {
		log.Printf("Failed to write to config file: %v", err)
		return
	}
}
