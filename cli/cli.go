package cli

import (
	"fmt"
	"github.com/kontena/kontena-client-go/client"
	"io/ioutil"
)

type CLI struct {
	Options Options
	Config  client.Config

	Client *client.Client
}

func (cli *CLI) Setup() {
	setupLogging(cli.Options)

	cli.Config.Logger = makeLogger(cli.Options, "client")
}

func (cli *CLI) Connect() error {
	if cli.Options.SSLCertPath == "" {

	} else if certPEM, err := ioutil.ReadFile(cli.Options.SSLCertPath); err != nil {
		return fmt.Errorf("Invalid --ssl-cert-path=%s: %v", cli.Options.SSLCertPath, err)
	} else {
		cli.Config.SSLCertPEM = certPEM
	}

	if client, err := cli.Config.Connect(); err != nil {
		return err
	} else {
		cli.Client = client
	}

	return nil
}
