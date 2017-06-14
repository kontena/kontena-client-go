package client

import (
	"encoding/json"
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

	client.Grids = gridsClient{client}

	return nil
}

func (client *Client) String() string {
	return fmt.Sprintf("%v", client.config.URL)
}

func (client *Client) Config() Config {
	return client.config
}

func (client *Client) doRequest(httpRequest *http.Request, responseBody interface{}) error {
	var httpResponse, err = client.httpClient.Do(httpRequest)

	if err != nil {
		return fmt.Errorf("HTTP %v %v request error: %v", httpRequest.Method, httpRequest.URL, err)
	} else {
		defer httpResponse.Body.Close()
	}

	var responseError = Error{
		httpRequest:  httpRequest,
		httpResponse: httpResponse,
	}

	if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {

	} else {
		responseBody = &responseError.apiError
	}

	if responseBody != nil && httpResponse.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(httpResponse.Body).Decode(responseBody); err != nil {
			return fmt.Errorf("HTTP %v %v response invalid: %v", httpRequest.Method, httpRequest.URL, err)
		}
	}

	log.Printf("[DEBUG] %v %v => HTTP %v %v: %#v",
		httpRequest.Method, httpRequest.URL,
		httpResponse.StatusCode, httpResponse.Status,
		responseBody,
	)

	switch httpResponse.StatusCode {
	case 200, 201:
		return nil
	case 404:
		return NotFoundError(responseError)
	default:
		return responseError
	}
}

func (client *Client) do(sling *sling.Sling, responseBody interface{}) error {
	if httpRequest, err := sling.Request(); err != nil {
		return fmt.Errorf("Invalid request: %v", err)
	} else if err := client.doRequest(httpRequest, responseBody); err != nil {
		return err
	} else {
		return nil
	}
}

func (client *Client) Get(path string, result interface{}) error {
	return client.do(client.sling.New().Get(path), result)
}
