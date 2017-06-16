package main

import (
	"github.com/kontena/kontena-client-go/cli/commands"

	"github.com/urfave/cli"
)

var kontenaCliGrids = commands.GridsCommand{
	CLI: &kontenaCli,
}

var gridsListCommand = cli.Command{
	Name:  "list",
	Usage: "List Grids",
	Action: func(c *cli.Context) error {
		return kontenaCliGrids.List()
	},
}

var gridsShowCommand = cli.Command{
	Name:      "show",
	Usage:     "Show Grid",
	ArgsUsage: "<grid>...",
	Action: func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return cli.NewExitError("Usage: kontena-cli grid show <grid>", 1)
		}

		for _, arg := range c.Args() {
			if err := kontenaCliGrids.Show(arg); err != nil {
				return err
			}
		}

		return nil
	},
}

var gridsCommand = cli.Command{
	Name:  "grid",
	Usage: "Kontena Grids",
	Subcommands: []cli.Command{
		gridsListCommand,
		gridsShowCommand,
	},
}
