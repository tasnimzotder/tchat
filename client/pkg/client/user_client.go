package client

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/tasnimzotder/tchat/_client/pkg/models"
)

func (c *Client) CreateUser(name, password string) (models.User, error) {
	body := models.User{
		Name:     name,
		Password: password,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return models.User{}, err
	}

	// Make a POST request to the server
	u := url.URL{
		Scheme: c.scheme,
		Host:   c.baseURL,
		Path:   "/v1/user/create",
	}

	resp, err := c.httpClient.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return models.User{}, err
	}

	defer resp.Body.Close()

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return models.User{}, err
	}

	// todo: remove this (temporary)
	// if password is empty, then user creation failed
	if user.Password == "" {
		user.Password = password
	}

	return user, nil
}
