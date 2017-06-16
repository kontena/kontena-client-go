package main

import (
	"log"
	"os"

	"github.com/kontena/kontena-client-go/client"
	"github.com/urfave/cli"
)

var config client.Config
var options struct {
	Token string
}

var globalClient *client.Client

func beforeApp(c *cli.Context) error {
	if clientToken, err := client.MakeToken(options.Token); err != nil {
		return err
	} else {
		config.Token = clientToken
	}

	log.Printf("[DEBUG] config: %#v", config)

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
	app().Run(os.Args)
}
