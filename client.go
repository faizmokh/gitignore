package main

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second,
		},
	}
}

func (c *Client) GetList() ([]List, error) {
	return []List{}, nil
}
