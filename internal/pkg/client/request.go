package client

import "net/http"

type RequestClient struct {
	runMode string
	client  *http.Client
}

func NewRequestClient(client *http.Client) *RequestClient {
	return &RequestClient{
		client: client,
	}
}
