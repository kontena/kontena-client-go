package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	config     Config
	apiURL     *url.URL
	httpClient *http.Client

	Grids GridsAPI
}

func (client *Client) init(config Config) error {
	client.config = config
	client.httpClient = config.httpClient()

	if apiURL, err := config.makeURL(); err != nil {
		return fmt.Errorf("Invalid URL: %v", err)
	} else {
		client.apiURL = apiURL
	}

	client.Grids = gridsClient{client}

	return nil
}

func (client *Client) String() string {
	return fmt.Sprintf("%v", client.config.URL)
}

func (client *Client) Config() Config {
	return client.config
}

func (client *Client) url(path ...string) *url.URL {
	if url, err := client.apiURL.Parse(strings.Join(path, "/")); err != nil {
		panic(err)
	} else {
		return url
	}
}

type request struct {
	Method       string
	URL          *url.URL
	RequestBody  interface{} // JSON
	ResponseBody interface{} // JSON
}

func (request request) String() string {
	return fmt.Sprintf("HTTP %v %v", request.Method, request.URL)
}

func (request request) encodeRequest() (*http.Request, error) {
	var requestBody io.Reader

	if request.RequestBody != nil {
		var requestBuffer bytes.Buffer

		if err := json.NewEncoder(&requestBuffer).Encode(request.RequestBody); err != nil {
			return nil, fmt.Errorf("Invalid request JSON: %v", err)
		}

		requestBody = &requestBuffer
	}

	var httpRequest, err = http.NewRequest(request.Method, request.URL.String(), requestBody)
	if err != nil {
		return nil, fmt.Errorf("Invalid request parameters: %v", err)
	}

	if requestBody != nil {
		httpRequest.Header.Add("Content-Type", "application/json")
	}

	return httpRequest, nil
}

func (request request) decodeResponse(httpRequest *http.Request, httpResponse *http.Response) error {
	var responseBody = request.ResponseBody
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
			return fmt.Errorf("Invalid response JSON: %v", err)
		}
	}

	switch httpResponse.StatusCode {
	case 200, 201:
		return nil
	case 404:
		return NotFoundError(responseError)
	default:
		return responseError
	}
}

func (client *Client) request(request request) error {
	if httpRequest, err := request.encodeRequest(); err != nil {
		return fmt.Errorf("Request %v invalid: %v", request, err)
	} else if httpResponse, err := client.httpClient.Do(httpRequest); err != nil {
		return fmt.Errorf("Request %v error: %v", request, err)
	} else {
		defer httpResponse.Body.Close()

		log.Printf("[DEBUG] %v %v => HTTP %v",
			httpRequest.Method, httpRequest.URL,
			httpResponse.Status,
		)

		return request.decodeResponse(httpRequest, httpResponse)
	}
}

func (client *Client) get(request request, path ...string) error {
	request.Method = "GET"
	request.URL = client.url(path...)

	return client.request(request)
}

func (client *Client) post(request request, path ...string) error {
	request.Method = "POST"
	request.URL = client.url(path...)

	return client.request(request)
}

func (client *Client) put(request request, path ...string) error {
	request.Method = "PUT"
	request.URL = client.url(path...)

	return client.request(request)
}

func (client *Client) delete(request request, path ...string) error {
	request.Method = "DELETE"
	request.URL = client.url(path...)

	return client.request(request)
}