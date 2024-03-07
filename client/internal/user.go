package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/pkg/config"
	"github.com/tasnimzotder/tchat/client/pkg/crypto"
	"github.com/tasnimzotder/tchat/client/pkg/file"
)

func InitializeNewConnection(name string) {
	body := models.User{
		Name: name,
	}

	// convert to io.Reader
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal body: %v", err)
		return
	}

	u := url.URL{
		Scheme: "http",
		Host:   config.GetEnvVariable("SERVER_HOST"),
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

	var encryption crypto.Encryptioner = &crypto.RSAEncryption{}

	// handle RSA key
	privateKey, publicKey, err := encryption.GenerateKeyPair(2048)
	if err != nil {
		log.Printf("Failed to generate key pair: %v", err)
		return
	}

	err = file.StoreRSAKeys(privateKey, publicKey)
	if err != nil {
		log.Printf("Failed to store RSA keys: %v", err)
		return
	}

	// save user id to config
	err = file.WriteToConfigFile(models.Config{
		ID:   response.ID,
		Name: response.Name,
	})

	if err != nil {
		log.Printf("Failed to write to config file: %v", err)
		return
	}
}
