package main

import (
	"os"

	kontena_cli "github.com/kontena/kontena-client-go/cli"
	"github.com/kontena/kontena-client-go/client"
	"github.com/urfave/cli"
)

var kontenaCli kontena_cli.CLI
var log = kontena_cli.Logger()

func beforeApp(c *cli.Context) error {
	kontenaCli.Setup()

	if clientToken, err := client.MakeToken(kontenaCli.Options.Token); err != nil {
		return err
	} else {
		kontenaCli.Config.Token = clientToken
	}

	log.Debugf("Config: %#v", kontenaCli.Config)

	if err := kontenaCli.Connect(); err != nil {
		return err
	}

	return nil
}

func app() *cli.App {
	app := cli.NewApp()
	app.Name = "kontena-cli"
	app.Usage = "Kontena CLI"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			EnvVar:      "KONTENA_DEBUG",
			Destination: &kontenaCli.Options.Debug,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Destination: &kontenaCli.Options.Verbose,
		},
		cli.BoolFlag{
			Name:        "quiet",
			Destination: &kontenaCli.Options.Quiet,
		},
		cli.StringFlag{
			Name:        "url",
			EnvVar:      "KONTENA_URL",
			Destination: &kontenaCli.Config.URL,
		},
		cli.StringFlag{
			Name:        "ssl-cert-cn",
			EnvVar:      "KONTENA_CERT_CN",
			Destination: &kontenaCli.Config.SSLServerName,
		},
		cli.StringFlag{
			Name:        "ssl-cert-file",
			EnvVar:      "KONTENA_CERT_FILE",
			Destination: &kontenaCli.Options.SSLCertPath,
		},
		cli.StringFlag{
			Name:        "token",
			EnvVar:      "KONTENA_TOKEN",
			Destination: &kontenaCli.Options.Token,
		},
	}
	app.Before = beforeApp
	app.Commands = []cli.Command{
		whoamiCommand,
		gridsCommand,
		nodesCommand,
	}

	return app
}

func main() {
	if err := app().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
