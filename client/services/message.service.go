package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/utils"
)

func GetMessages() ([]models.Message, error) {
	config, err := utils.ReadFromConfigFile()
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
		return nil, err
	}

	u := url.URL{
		Scheme: "http",
		Host:   utils.GetEnvVariable("SERVER_HOST"),
		Path:   "/v1/message/get",
	}

	params := url.Values{}

	params.Add("id", config.ID)
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//if response status code is not 200
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var messages []models.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func SendMessage(recipientID, messageType, message string) error {
	config, err := utils.ReadFromConfigFile()
	if err != nil {
		return err
	}

	u := url.URL{
		Scheme: "http",
		Host:   utils.GetEnvVariable("SERVER_HOST"),
		Path:   "/v1/message/send",
	}

	body := models.Message{
		SenderID:    config.ID,
		RecipientID: recipientID,
		Payload:     message,
		MessageType: messageType,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewReader(bodyBytes),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	//if response status code is not 201
	if resp.StatusCode != http.StatusCreated {
		return err
	}

	return nil
}
