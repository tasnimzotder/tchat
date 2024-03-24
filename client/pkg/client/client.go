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
		baseURL:    baseURL,
		scheme:     scheme,
		httpClient: http.Client{
			// Timeout: time.Second * 10,
		},
	}
}
