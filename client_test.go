package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	okResponse = `{
			"git": {
				"name": "Git",
				"key": "git",
				"fileName": "Git.gitignore",
				"contents": "contents"
			}
		}`
)

func TestClientGetList(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(okResponse))
	})

	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewClient()
	cli.httpClient = httpClient

	list, _ := cli.GetList()

	if len(list) != 0 {
		t.Errorf("list should not be empty, got: %d, want %d", len(list), 1)
	}
}

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}
