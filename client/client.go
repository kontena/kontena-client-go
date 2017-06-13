package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/sling"
)

type Client struct {
	config     Config
	httpClient *http.Client
	sling      *sling.Sling

	Grids GridsAPI
}

func (client *Client) init(config Config) error {
	client.config = config
	client.httpClient = config.httpClient()
	client.sling = sling.New().Client(client.httpClient).Base(config.URL)

	client.Grids = gridsClient{client.sling}

	return nil
}

func (client *Client) String() string {
	return fmt.Sprintf("%v", client.config.URL)
}

func (client *Client) Config() Config {
	return client.config
}

func (client *Client) Get(path string, result interface{}) error {
	return do(client.sling.New().Get(path), result)
}

func do(sling *sling.Sling, responseBody interface{}) error {
	var clientError Error

	if httpRequest, err := sling.Request(); err != nil {
		return fmt.Errorf("Invalid request: %v", err)
	} else if httpResponse, err := sling.Do(httpRequest, responseBody, &clientError.API); err != nil {
		return fmt.Errorf("%v %v: %v", httpRequest.Method, httpRequest.URL, err)
	} else {
		clientError.httpRequest = httpRequest
		clientError.httpResponse = httpResponse

		log.Printf("[DEBUG] %v %v => HTTP %v %v: %#v",
			httpRequest.Method, httpRequest.URL,
			httpResponse.StatusCode, httpResponse.Status,
			responseBody,
		)

		switch httpResponse.StatusCode {
		case 200, 201:
			return nil
		case 404:
			return NotFoundError(clientError)
		default:
			return clientError
		}
	}
}
