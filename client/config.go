package client

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

type Config struct {
	URL         string
	AccessToken string
}

func (config Config) makeURL() (*url.URL, error) {
	return url.Parse(config.URL)
}

func (config Config) httpClient() *http.Client {
	var oauthConfig = &oauth2.Config{}
	var oauthToken = &oauth2.Token{
		AccessToken: config.AccessToken,
	}

	return oauthConfig.Client(oauth2.NoContext, oauthToken)
}

func (config Config) MakeClient() (*Client, error) {
	var client Client

	if config.AccessToken == "" {
		return nil, fmt.Errorf("AccessToken is required")
	}

	if err := client.init(config); err != nil {
		return nil, err
	}

	return &client, nil
}
