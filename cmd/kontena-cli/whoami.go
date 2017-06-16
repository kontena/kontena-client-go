package main

import (
	"github.com/urfave/cli"

	"github.com/kontena/kontena-client-go/cli/commands"
)

var kontenaCliWhoami = commands.WhoamiCommand{
	CLI: &kontenaCli,
}

var whoamiCommand = cli.Command{
	Name:  "whoami",
	Usage: "Show kontena master user",
	Action: func(c *cli.Context) error {
		return kontenaCliWhoami.Whoami()
	},
}
