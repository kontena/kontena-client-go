package client

import (
	"github.com/kontena/terraform-provider-kontena/api"
)

func (client *Client) Ping() error {
	var ping api.Ping

	return client.get(request{ResponseBody: &ping}, "/v1/ping")
}
