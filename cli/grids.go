package cli

import (
	"github.com/kontena/kontena-client-go/api"
)

func printGrids(grids []api.Grid) {
	tbl := makeTable("Name", "Nodes", "Services", "Users")

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
	*CLI
}

func (cmd GridsCommand) List() error {
	if grids, err := cmd.client.Grids.List(); err != nil {
		return err
	} else {
		printGrids(grids)
	}

	return nil
}

func (cmd GridsCommand) Show(name string) error {
	if grid, err := cmd.client.Grids.Get(name); err != nil {
		return err
	} else if err := print(grid); err != nil {
		return err
	}

	return nil
}
