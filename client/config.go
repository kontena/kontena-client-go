package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

type Config struct {
	URL           string
	SSLCertPEM    []byte
	SSLServerName string
	ClientID      string // default OAuth2ClientID
	ClientSecret  string // default OAuth2ClientSecret
	Token         *Token // default does anonymous requests without any access token
	Logger        Logger
}

func (config Config) isSSL() bool {
	if configURL, err := url.Parse(config.URL); err != nil {
		return false
	} else if configURL.Scheme == "https" {
		return true
	} else if configURL.Scheme == "http" {
		return false
	}

	return false
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

func (config Config) tlsConfig() (*tls.Config, error) {
	var tlsConfig *tls.Config

	if config.isSSL() {
		tlsConfig = &tls.Config{
			ServerName: config.SSLServerName,
		}

		if config.SSLCertPEM != nil {
			var certPool = x509.NewCertPool()

			if ok := certPool.AppendCertsFromPEM(config.SSLCertPEM); !ok {
				return nil, fmt.Errorf("Invalid config.SSLCertPEM")
			}

			tlsConfig.RootCAs = certPool
		}
	}

	return tlsConfig, nil
}

// Create an http.Client for the server, using the tls configuration.
//
// This does not include the oauth2 access token.
func (config Config) httpClient() (*http.Client, error) {
	var httpTransport http.RoundTripper

	if tlsConfig, err := config.tlsConfig(); err != nil {
		return nil, err
	} else {
		httpTransport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	var httpClient = &http.Client{
		Transport: httpTransport,
	}

	return httpClient, nil
}

func (config Config) oauthContext() (context.Context, error) {
	var ctx = context.Background()

	if httpClient, err := config.httpClient(); err != nil {
		return nil, err
	} else {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	}

	return ctx, nil
}

func (config Config) oauthConfig() (*oauth2.Config, error) {
	var oauthConfig = oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
	}

	// apply defaults
	if oauthConfig.ClientID == "" {
		oauthConfig.ClientID = OAuth2ClientID
	}
	if oauthConfig.ClientSecret == "" {
		oauthConfig.ClientSecret = OAuth2ClientSecret
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

// ExchangeToken allows you to exchange a single-use oauth2 code for an access
// token.
//
// This does not need to have any config.Token set.
func (config Config) ExchangeToken(code string) (*Token, error) {
	if oauthContext, err := config.oauthContext(); err != nil {
		return nil, err
	} else if oauthConfig, err := config.oauthConfig(); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 config: %v", err)
	} else if oauthToken, err := oauthConfig.Exchange(oauthContext, code); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 code: %v", err)
	} else {
		return (*Token)(oauthToken), nil
	}
}

// Create an http.Client using the OAuth2 configuration, using the oauth2 access
// token in requests.
//
// XXX: if the token expires, then oauth2 refreshes it, and the caller needs to
// persist that...?
func (config Config) oauthClient() (*http.Client, error) {
	if config.Token == nil {
		// no oauth2 token
		return config.httpClient()
	}

	if oauthContext, err := config.oauthContext(); err != nil {
		return nil, err
	} else if oauthConfig, err := config.oauthConfig(); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 config: %v", err)
	} else {
		var oauthToken = (*oauth2.Token)(config.Token)
		var httpClient = oauthConfig.Client(oauthContext, oauthToken)

		return httpClient, nil
	}
}
