package client

import "github.com/kontena/kontena-client-go/api"

type UsersAPI interface {
	GetUser() (api.User, error)
}

type usersClient struct {
	client *Client
}

func (usersClient usersClient) GetUser() (api.User, error) {
	var user api.User

	return user, usersClient.client.get(request{ResponseBody: &user}, "/v1/user")
}
