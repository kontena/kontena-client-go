package commands

import (
	"github.com/kontena/kontena-client-go/api"
	"github.com/kontena/kontena-client-go/cli"
)

func printGrids(grids []api.Grid) {
	tbl := cli.Table("Name", "Nodes", "Services", "Users")

	for _, grid := range grids {
		tbl.AddRow(grid.Name,
			grid.NodeCount,
			grid.ServiceCount,
			grid.UserCount,
		)
	}

	tbl.Print()
}

type GridsCommand struct {
	*cli.CLI
}

// List prints a table of all grids.
func (cmd GridsCommand) List() error {
	if grids, err := cmd.Client.Grids.List(); err != nil {
		return err
	} else {
		printGrids(grids)
	}

	return nil
}

// Show prints details of a given grid.
func (cmd GridsCommand) Show(name string) error {
	if grid, err := cmd.Client.Grids.Get(name); err != nil {
		return err
	} else {
		return cli.Print(grid)
	}
}
