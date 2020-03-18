package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	okResponse = `{
		"goodsync": {
			 "name": "GoodSync",
			 "key": "goodsync",
			 "fileName": "GoodSync.gitignore",
			 "contents": "\n### GoodSync ###\n_gsdata_\n"
		 },
		 "psoccreator": {
			 "name": "PSoCCreator",
			 "key": "psoccreator",
			 "fileName": "PSoCCreator.gitignore",
			 "contents": "\n### PSoCCreator ###\n# Project Settings\n*.cywrk.*\n*.cyprj.*\n\n# Generated Assets and Resources\nDebug/\nRelease/\nExport/\n*/codegentemp\n*/Generated_Source\n*_datasheet.pdf\n*_timing.html\n*.cycdx\n*.cyfit\n*.rpt\n*.svd\n*.log\n*.zip\n"
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

	list, err := cli.GetList()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(list), "it should not be empty")
	assert.Equal(t, "GoodSync", list["goodsync"].Name, "it should match to git")
}

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

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
