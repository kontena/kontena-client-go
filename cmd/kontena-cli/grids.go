package main

import (
  "github.com/urfave/cli"
  "github.com/rodaine/table"
  "github.com/kontena/kontena-client-go/api"
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
	Name:   "list",
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

var gridsCommand = cli.Command{
	Name:  "grids",
	Usage: "Kontena Grids",
  Subcommands: []cli.Command{
    gridsListCommand,
  },
}
