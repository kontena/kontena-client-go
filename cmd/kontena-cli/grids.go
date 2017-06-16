package main

import (
	"github.com/kontena/kontena-client-go/api"
	"github.com/rodaine/table"
	"github.com/urfave/cli"
)

func printGrids(grids []api.Grid) {
	tbl := table.New("Name", "Nodes", "Services", "Users")

	for _, grid := range grids {
		tbl.AddRow(grid.Name,
			grid.NodeCount,
			grid.ServiceCount,
			grid.UserCount,
		)
	}

	tbl.Print()
}

var gridsListCommand = cli.Command{
	Name:  "list",
	Usage: "List Grids",
	Action: func(c *cli.Context) error {
		if grids, err := globalClient.Grids.List(); err != nil {
			return err
		} else {
			printGrids(grids)
		}

		return nil
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
			if grid, err := globalClient.Grids.Get(arg); err != nil {
				return err
			} else if err := print(grid); err != nil {
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
