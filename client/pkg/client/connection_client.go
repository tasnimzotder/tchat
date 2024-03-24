package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/url"
)

type ConnectionRequest struct {
	Name        string `json:"name"`
	UserID      string `json:"user_id"`
	PublicKey   string `json:"public_key"`
	PassKey     string `json:"pass_key"`
	Expiration  string `json:"expiration"`
	Persistence string `json:"persistent"`
}

func (c *Client) CreateConnection(req ConnectionRequest) error {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	u := url.URL{
		Scheme: c.scheme,
		Host:   c.baseURL,
		Path:   "/v1/connection/set",
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

	if res.StatusCode != 200 {
		return errors.New("failed to create connection")
	}

	return nil
}

type GetConnectionRequest struct {
	UserID  string `json:"user_id"`
	PassKey string `json:"pass_key"`
}

// type GetConnectionResponse = ConnectionRequest

func (c *Client) GetConnection(req GetConnectionRequest) (ConnectionRequest, error) {
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return ConnectionRequest{}, err
	}

	u := url.URL{
		Scheme: c.scheme,
		Host:   c.baseURL,
		Path:   "/v1/connection/get",
	}

	res, err := c.httpClient.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return ConnectionRequest{}, err
	}

	defer res.Body.Close()

	var connection ConnectionRequest
	if err := json.NewDecoder(res.Body).Decode(&connection); err != nil {
		return ConnectionRequest{}, err
	}

	return connection, nil
}
