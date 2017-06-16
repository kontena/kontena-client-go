package client

import (
	"github.com/kontena/kontena-client-go/api"
)

func (client *Client) Ping() error {
	var ping api.Ping

	return client.get(request{ResponseBody: &ping}, "/v1/ping")
}
