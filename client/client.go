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

	client.Grids = gridsClient{client.sling.New().Path("/v1/grids/")}

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
	var requestError Error

	if httpRequest, err := sling.Request(); err != nil {
		return fmt.Errorf("Invalid request: %v", err)
	} else if httpResponse, err := sling.Do(httpRequest, responseBody, &requestError.API); err != nil {
		return fmt.Errorf("%v %v: %v", httpRequest.Method, httpRequest.URL, err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		requestError.httpRequest = httpRequest
		requestError.httpResponse = httpResponse

		return requestError
	} else {
		log.Printf("[DEBUG] %v %v => HTTP %v %v: %#v",
			httpRequest.Method, httpRequest.URL,
			httpResponse.StatusCode, httpResponse.Status,
			responseBody,
		)

		return nil
	}
}
