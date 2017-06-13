package client

import (
	"github.com/kontena/terraform-provider-kontena/api"
)

func (client *Client) Ping() error {
	var ping api.Ping

	return client.Get("/v1/ping", &ping)
}
