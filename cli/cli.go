package cli

import (
	"github.com/kontena/kontena-client-go/client"
)

type CLI struct {
	Options Options
	Config  client.Config

	client *client.Client
}

func (cli *CLI) Connect() error {
	if client, err := cli.Config.Connect(); err != nil {
		return err
	} else {
		cli.client = client
	}

	return nil
}
