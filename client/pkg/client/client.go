package client

import (
	"net/http"
)

type Client struct {
	baseURL    string
	scheme     string
	httpClient http.Client
}

func NewClient(baseURL, scheme string) *Client {
	return &Client{
		baseURL: baseURL,
		scheme:  scheme,
		httpClient: http.Client{
			// Timeout: time.Second * 10,
		},
	}
}

// func (c *Client) GetMessages() ([]models.Message, error) {
// 	// Make a GET request to the server
// 	resp, err := http.Get(c.baseURL + "/messages")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// ...
// 	return nil, nil
// }
