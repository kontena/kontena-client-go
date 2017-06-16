package main

import (
	"os"

	"github.com/kontena/kontena-client-go/client"
	"github.com/op/go-logging"
	"github.com/urfave/cli"
)

var config client.Config
var options struct {
	Debug   bool
	Verbose bool
	Quiet   bool
	Token   string
}

var globalClient *client.Client

func beforeApp(c *cli.Context) error {
	if options.Debug {
		logging.SetLevel(logging.DEBUG, "kontena-cli")
	} else if options.Verbose {
		logging.SetLevel(logging.INFO, "kontena-cli")
	} else if options.Quiet {
		logging.SetLevel(logging.ERROR, "kontena-cli")
	} else {
		logging.SetLevel(logging.WARNING, "kontena-cli")
	}

	if clientToken, err := client.MakeToken(options.Token); err != nil {
		return err
	} else {
		config.Token = clientToken
	}

	log.Debugf("app config: %#v", config)

	if client, err := config.Connect(); err != nil {
		return err
	} else {
		globalClient = client
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
			Destination: &options.Debug,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Destination: &options.Verbose,
		},
		cli.BoolFlag{
			Name:        "quiet",
			Destination: &options.Quiet,
		},
		cli.StringFlag{
			Name:        "url",
			EnvVar:      "KONTENA_URL",
			Destination: &config.URL,
		},
		cli.StringFlag{
			Name:        "token",
			EnvVar:      "KONTENA_TOKEN",
			Destination: &options.Token,
		},
	}
	app.Before = beforeApp
	app.Commands = []cli.Command{
		whoamiCommand,
		gridsCommand,
	}

	return app
}

func main() {
	if err := app().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
