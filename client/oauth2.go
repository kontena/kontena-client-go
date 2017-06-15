package client

import "golang.org/x/oauth2"

const (
	// generate using many fair dice rolls
	OAUTH2_CLIENT_ID     = "d0f032e5af41187fdaf45b4aeee76ee4"
	OAUTH2_CLIENT_SECRET = "4141f52bb9f80dfc776b5b773ccf550d"
)

type Token oauth2.Token

func MakeToken(accessToken string) *Token {
	return (*Token)(&oauth2.Token{
		AccessToken: accessToken,
	})
}
