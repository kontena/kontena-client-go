package main

import (
	"fmt"

	kontena_cli "github.com/kontena/kontena-client-go/cli"
	"github.com/urfave/cli"
)

var kontenaCliNodes = kontena_cli.NodesCommand{
	CLI: &kontenaCli,
}

var nodesListCommand = cli.Command{
	Name:  "list",
	Usage: "List Nodes",
	Action: func(c *cli.Context) error {
		return kontenaCliNodes.List()
	},
}

var nodesCommand = cli.Command{
	Name:  "node",
	Usage: "Kontena Nodes",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "grid",
			EnvVar:      "KONTENA_GRID",
			Destination: &kontenaCliNodes.Grid,
		},
	},
	Before: func(c *cli.Context) error {
		if kontenaCliNodes.Grid == "" {
			return fmt.Errorf("Missing --grid")
		}

		return nil
	},
	Subcommands: []cli.Command{
		nodesListCommand,
	},
}
