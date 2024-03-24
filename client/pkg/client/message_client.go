package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/tasnimzotder/tchat/_client/pkg/models"
)

func (c *Client) SendMessage(message models.Message) error {
	u := url.URL{
		Scheme: c.scheme,
		Host:   c.baseURL,
		Path:   "/v1/message/send",
	}

	bodyBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 201 {
		return errors.New("failed to send message")
	}

	return nil
}

func (c *Client) GetMessages(id string) ([]models.Message, error) {
	u := url.URL{
		Scheme: c.scheme,
		Host:   c.baseURL,
		Path:   "/v1/message/get",
	}

	params := url.Values{}
	params.Add("id", id)

	u.RawQuery = params.Encode()

	res, err := c.httpClient.Get(
		u.String(),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// statusRegex := regexp.MustCompile(`^2\d{2}$`)
	// if !statusRegex.MatchString(res.Status) {
	// 	return nil, errors.New("failed to get messages")
	// }

	var messages []models.Message
	_ = json.NewDecoder(res.Body).Decode(&messages)

	// if err != nil {
	// 	return nil, err
	// }

	return messages, nil
}
