package services

import (
	"bytes"
	"encoding/json"
	"github.com/tasnimzotder/tchat/client/models"
	"github.com/tasnimzotder/tchat/client/utils"
	"log"
	"net/http"
	"net/url"
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
		log.Printf("Failed to get messages: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	//if response status code is not 200
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get messages: %v", resp.StatusCode)
		return nil, err
	}

	var messages []models.Message
	err = json.NewDecoder(resp.Body).Decode(&messages)

	if err != nil {
		log.Printf("Failed to decode messages: %v", err)
		return nil, err
	}

	return messages, nil
}

func SendMessage(recipientID, messageType, message string) error {
	config, err := utils.ReadFromConfigFile()
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
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
		log.Printf("Failed to marshal message: %v", err)
		return err
	}

	resp, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewReader(bodyBytes),
	)

	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	defer resp.Body.Close()

	//if response status code is not 201
	if resp.StatusCode != http.StatusCreated {
		log.Printf("Failed to send message: %v", resp.StatusCode)
		return err
	}

	return nil
}
