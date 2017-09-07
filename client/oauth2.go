package client

import (
	"fmt"

	"golang.org/x/oauth2"
)

const (
	// generate using many fair dice rolls
	Oauth2ClientID     = "d0f032e5af41187fdaf45b4aeee76ee4"
	Oauth2ClientSecret = "4141f52bb9f80dfc776b5b773ccf550d"
)

type Token oauth2.Token

func MakeToken(accessToken string) (*Token, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("Empty oauth2 token")
	}

	var oauthToken = &oauth2.Token{
		AccessToken: accessToken,
	}

	return (*Token)(oauthToken), nil
}
