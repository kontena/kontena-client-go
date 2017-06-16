package main

import (
	"fmt"

	"github.com/kontena/kontena-client-go/cli/commands"
	"github.com/urfave/cli"
)

var kontenaCliNodes = commands.NodesCommand{
	CLI: &kontenaCli,
}

var nodesListCommand = cli.Command{
	Name:  "list",
	Usage: "List Nodes",
	Action: func(c *cli.Context) error {
		return kontenaCliNodes.List()
	},
}

var nodesShowCommand = cli.Command{
	Name:      "show",
	Usage:     "Show Nodes",
	ArgsUsage: "<node>...",
	Action: func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return cli.NewExitError("Usage: kontena-cli node show <name>", 1)
		}

		for _, arg := range c.Args() {
			if err := kontenaCliNodes.Show(arg); err != nil {
				return err
			}
		}

		return nil
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
		nodesShowCommand,
	},
}
