package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MakeClient calls ConnectOAuth2() to update the config.LoginToken.
func (config Config) MakeClient() (*Client, error) {
	var client = Client{
		logger: config.Logger,
		config: config,
	}

	if apiURL, err := config.makeURL(); err != nil {
		return nil, fmt.Errorf("Invalid API URL %v: %v", config.URL, err)
	} else {
		client.apiURL = apiURL
	}

	if httpClient, err := config.oauthClient(); err != nil {
		return nil, err
	} else {
		client.httpClient = httpClient
	}

	if err := client.init(); err != nil {
		return nil, err
	}

	return &client, nil
}

// Connect uses config information to allow the client to make a connection.
func (config Config) Connect() (*Client, error) {
	if client, err := config.MakeClient(); err != nil {
		return nil, err
	} else if err := client.Ping(); err != nil {
		return nil, err
	} else {
		return client, nil
	}
}

// Client contains exported fields that give access to parts of the API to then
// make calls on.
type Client struct {
	logger     Logger
	config     Config
	httpClient *http.Client
	apiURL     *url.URL

	Users              UsersAPI
	Grids              GridsAPI
	Nodes              NodesAPI
	ExternalRegistries ExternalRegistriesAPI
}

func (client *Client) init() error {
	client.Users = usersClient{client}
	client.Grids = gridsClient{client}
	client.Nodes = nodesClient{client}
	client.ExternalRegistries = externalRegistriesClient{client}

	return nil
}

func (client *Client) String() string {
	return fmt.Sprintf("%v", client.config.URL)
}

// Config returns the current configurations that have been setup with the
// client (config.go).
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

// String returns the request method and URL as a string.
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
	case 403:
		return ForbiddenError(responseError)
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

		if client.logger != nil {
			client.logger.Infof("%v %v => HTTP %v",
				httpRequest.Method, httpRequest.URL,
				httpResponse.Status,
			)
		}

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
