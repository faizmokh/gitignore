package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (c *Client) GetList() (List, error) {
	req, err := http.NewRequest("GET", "https://www.gitignore.io/api/list?format=json", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var list List
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&list); err != nil {
		return nil, err
	}

	return list, nil
}
