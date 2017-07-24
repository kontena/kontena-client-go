package cli

import (
	"github.com/kontena/kontena-client-go/client"
)

type CLI struct {
	Options Options
	Config  client.Config

	Client *client.Client
}

func (cli *CLI) Setup() {
	setupLogging(cli.Options)
}

func (cli *CLI) Connect() error {
	cli.Config.Logger = makeLogger(cli.Options, "client")

	if client, err := cli.Config.Connect(); err != nil {
		return err
	} else {
		cli.Client = client
	}

	return nil
}
