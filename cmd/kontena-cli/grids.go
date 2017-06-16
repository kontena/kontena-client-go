package main

import (
	"github.com/kontena/kontena-client-go/api"
	"github.com/rodaine/table"
	"github.com/urfave/cli"
)

var gridFlag = cli.StringFlag{
	Name:        "grid",
	EnvVar:      "KONTENA_GRID",
	Destination: &Options.Grid,
}

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

func listGrids() error {
	if grids, err := Client.Grids.List(); err != nil {
		return err
	} else {
		printGrids(grids)
	}

	return nil
}

func showGrid(name string) error {
	if grid, err := Client.Grids.Get(name); err != nil {
		return err
	} else if err := print(grid); err != nil {
		return err
	}

	return nil
}

var gridsListCommand = cli.Command{
	Name:  "list",
	Usage: "List Grids",
	Action: func(c *cli.Context) error {
		return listGrids()
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
			if err := showGrid(arg); err != nil {
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
