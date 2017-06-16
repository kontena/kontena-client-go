package main

import (
	"github.com/urfave/cli"
)

var whoamiCommand = cli.Command{
	Name:  "whoami",
	Usage: "Show kontena master user",
	Action: func(c *cli.Context) error {
		return kontenaCli.Whoami()
	},
}
