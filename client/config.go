package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

type Config struct {
	URL          string
	ClientID     string // default OAUTH2_CLIENT_ID
	ClientSecret string // default OAUTH2_CLIENT_SECRET
	Token        *Token
}

func (config Config) makeURL(path ...string) (*url.URL, error) {
	if baseURL, err := url.Parse(config.URL); err != nil {
		return nil, err
	} else if pathURL, err := baseURL.Parse(strings.Join(path, "/")); err != nil {
		return nil, err
	} else {
		return pathURL, nil
	}
}

func (config Config) oauthConfig() (*oauth2.Config, error) {
	var oauthConfig = oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
	}

	// apply defaults
	if oauthConfig.ClientID == "" {
		oauthConfig.ClientID = OAUTH2_CLIENT_ID
	}
	if oauthConfig.ClientSecret == "" {
		oauthConfig.ClientSecret = OAUTH2_CLIENT_SECRET
	}

	// apply oauth2 API URLs
	if authURL, err := config.makeURL("/oauth2/authorize"); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 authorize URL: %v", err)
	} else {
		oauthConfig.Endpoint.AuthURL = authURL.String()
	}
	if tokenURL, err := config.makeURL("/oauth2/token"); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 token URL: %v", err)
	} else {
		oauthConfig.Endpoint.TokenURL = tokenURL.String()
	}

	return &oauthConfig, nil
}

// Exchange a single-use oauth2 code for an access token.
//
// This does not need to have any config.Token set.
func (config Config) ExchangeToken(code string) (*Token, error) {
	if oauthConfig, err := config.oauthConfig(); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 config: %v", err)
	} else if oauthToken, err := oauthConfig.Exchange(context.TODO(), code); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 code: %v", err)
	} else {
		return (*Token)(oauthToken), nil
	}
}

// Create an http.Client using the OAuth2 configuration, using the oauth2 access token in requests.
//
// This requires the config.Token to be set.
//
// XXX: if the token expires, then oauth2 refreshes it, and the caller needs to persist that...?
func (config Config) oauthClient() (*http.Client, error) {
	if oauthConfig, err := config.oauthConfig(); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 config: %v", err)
	} else if config.Token == nil {
		return nil, fmt.Errorf("Missing oauth2 token")
	} else {
		var httpClient = oauthConfig.Client(context.TODO(), (*oauth2.Token)(config.Token))

		return httpClient, nil
	}
}
