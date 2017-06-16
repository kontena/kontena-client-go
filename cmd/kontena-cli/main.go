package main

import (
	"os"

	"github.com/kontena/kontena-client-go/client"
	"github.com/op/go-logging"
	"github.com/urfave/cli"
)

var Config client.Config
var Options options
var Client *client.Client

func beforeApp(c *cli.Context) error {
	if Options.Debug {
		logging.SetLevel(logging.DEBUG, "kontena-cli")
	} else if Options.Verbose {
		logging.SetLevel(logging.INFO, "kontena-cli")
	} else if Options.Quiet {
		logging.SetLevel(logging.ERROR, "kontena-cli")
	} else {
		logging.SetLevel(logging.WARNING, "kontena-cli")
	}

	if clientToken, err := client.MakeToken(Options.Token); err != nil {
		return err
	} else {
		Config.Token = clientToken
	}

	log.Debugf("app config: %#v", Config)

	if client, err := Config.Connect(); err != nil {
		return err
	} else {
		Client = client
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
			Destination: &Options.Debug,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Destination: &Options.Verbose,
		},
		cli.BoolFlag{
			Name:        "quiet",
			Destination: &Options.Quiet,
		},
		cli.StringFlag{
			Name:        "url",
			EnvVar:      "KONTENA_URL",
			Destination: &Config.URL,
		},
		cli.StringFlag{
			Name:        "token",
			EnvVar:      "KONTENA_TOKEN",
			Destination: &Options.Token,
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
