package client

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type Client struct {
	config     Config
	httpClient *http.Client
	sling      *sling.Sling
}

func (client *Client) init(config Config) error {
	client.config = config
	client.httpClient = config.httpClient()
	client.sling = sling.New().Client(client.httpClient).Base(config.URL)

	return nil
}

func (client *Client) String() string {
	return fmt.Sprintf("%v", client.config.URL)
}

func (client *Client) Config() Config {
	return client.config
}

func (client *Client) Get(path string, result interface{}) error {
	var requestError Error

	if httpResponse, err := client.sling.New().Get(path).Receive(result, &requestError.API); err != nil {
		return err
	} else if httpResponse.StatusCode != 200 {
		requestError.HTTPStatus = httpResponse.StatusCode

		return requestError
	} else {
		return nil
	}
}
