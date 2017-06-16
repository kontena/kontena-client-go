package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kontena/kontena-client-go/client"
	"github.com/spf13/cobra"
)

var options struct {
	Token string
}

var cli struct {
	config client.Config
	client *client.Client
}

var rootCommand = &cobra.Command{
	Use: "kontena",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if clientToken, err := client.MakeToken(options.Token); err != nil {
			return err
		} else {
			cli.config.Token = clientToken
		}

		log.Printf("[DEBUG] config: %#v", cli.config)

		if client, err := cli.config.Connect(); err != nil {
			return err
		} else {
			cli.client = client
		}

		return nil
	},
}

func init() {
	rootCommand.PersistentFlags().StringVar(&cli.config.URL, "url", "", "HTTP URL to kontena-server API")
	rootCommand.PersistentFlags().StringVar(&options.Token, "token", "", "OAuth2 token")

	rootCommand.AddCommand(whoamiCommand)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
