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
	ClientID      string // default OAUTH2_CLIENT_ID
	ClientSecret  string // default OAUTH2_CLIENT_SECRET
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
	var (
		pathURL *url.URL
		baseURL *url.URL
		err error
	)

	if baseURL, err = url.Parse(config.URL); err != nil {
		return nil, err
	}

	if pathURL, err = baseURL.Parse(strings.Join(path, "/")); err != nil {
		return nil, err
	}

	return pathURL, err
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
	var (
		tlsConfig *tls.Config
		err error
	)

	if tlsConfig, err = config.tlsConfig(); err != nil {
		return nil, err
	}

	httpTransport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}

	return httpClient, nil
}

func (config Config) oauthContext() (context.Context, error) {
	var (
		httpClient *http.Client
		err error
	)

	if httpClient, err = config.httpClient(); err != nil {
		return nil, err
	}

	return context.WithValue(context.Background(),
		oauth2.HTTPClient, httpClient), err
}

func (config Config) oauthConfig() (*oauth2.Config, error) {
	oauthConfig := oauth2.Config{
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

	var (
		authURL *url.URL
		tokenURL *url.URL
		err error
	)

	// apply oauth2 API URLs
	if authURL, err = config.makeURL("/oauth2/authorize"); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 authorize URL: %v", err)
	}

	oauthConfig.Endpoint.AuthURL = authURL.String()

	if tokenURL, err = config.makeURL("/oauth2/token"); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 token URL: %v", err)
	}

	oauthConfig.Endpoint.TokenURL = tokenURL.String()

	return &oauthConfig, nil
}

// ExchangeToken allows you to exchange a single-use oauth2 code for an access
// token.
//
// This does not need to have any config.Token set.
func (config Config) ExchangeToken(code string) (*Token, error) {
	var (
		oauthToken *oauth2.Token
		oauthConfig *oauth2.Config
		oauthContext context.Context
		err error
	)

	if oauthContext, err = config.oauthContext(); err != nil {
		return nil, err
	} else if oauthConfig, err = config.oauthConfig(); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 config: %v", err)
	} else if oauthToken, err = oauthConfig.Exchange(oauthContext, code); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 code: %v", err)
	}

	return (*Token)(oauthToken), nil
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

	var (
		oauthConfig *oauth2.Config
		oauthContext context.Context
		err error
	)

	if oauthContext, err = config.oauthContext(); err != nil {
		return nil, err
	} else if oauthConfig, err = config.oauthConfig(); err != nil {
		return nil, fmt.Errorf("Invalid oauth2 config: %v", err)
	}

	var oauthToken = (*oauth2.Token)(config.Token)
	var httpClient = oauthConfig.Client(oauthContext, oauthToken)

	return httpClient, nil
}
