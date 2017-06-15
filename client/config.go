package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

type Config struct {
	URL          string
	ClientID     string // default OAUTH2_CLIENT_ID
	ClientSecret string // default OAUTH2_CLIENT_SECRET
	LoginCode    string // single-use oauth2 code, exchanged for an oauth2 token stored in LoginToken
	LoginToken   string // oauth2 token, type is always Bearer
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

// Create an http.Client using the OAuth2 configuration, using the oauth2 access token in requests.
//
// If an AccessToken is given, use the given token.
// If an InitialAdminCode is given, exchange it for an oauth2 token, and store that back in AccessToken.
//
// Modifies the *Config to update the AccessToken as necessary.
func (config *Config) ConnectOAuth2() (*http.Client, error) {
	var oauthConfig = &oauth2.Config{
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

	// resolve token
	var oauthToken *oauth2.Token

	if config.LoginToken != "" {
		// use given token
		oauthToken = &oauth2.Token{
			AccessToken: config.LoginToken,
		}
	} else if config.LoginCode != "" {
		if token, err := oauthConfig.Exchange(context.TODO(), config.LoginCode); err != nil {
			return nil, fmt.Errorf("Invalid oauth2 LoginCode: %v", err)
		} else {
			oauthToken = token

			log.Printf("[INFO] Exchange oauth2 code %v for oauth2 token: %#v", config.LoginCode, token)

			// update config
			config.LoginCode = "" // single-use
			config.LoginToken = token.AccessToken
		}
	} else {
		return nil, fmt.Errorf("No oauth2 LoginToken or LoginCode given")
	}

	var httpClient = oauthConfig.Client(context.TODO(), oauthToken)

	return httpClient, nil
}
