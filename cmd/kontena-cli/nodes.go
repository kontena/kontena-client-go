package main

import (
	"github.com/kontena/kontena-client-go/api"
	"github.com/rodaine/table"
	"github.com/urfave/cli"
)

func printNodes(nodes []api.Node) {
	tbl := table.New("Name", "Version", "Status", "Initial", "Labels")

	for _, node := range nodes {
		tbl.AddRow(node.Name)
	}

	tbl.Print()
}

func listNodes(grid string) error {
	if nodes, err := Client.Nodes.List(grid); err != nil {
		return err
	} else {
		printNodes(nodes)
	}

	return nil
}

var nodesListCommand = cli.Command{
	Name:  "list",
	Usage: "List Nodes",
	Action: func(c *cli.Context) error {
		if grid, err := Options.requireGrid(); err != nil {
			return err
		} else {
			return listNodes(grid)
		}
	},
}

var nodesCommand = cli.Command{
	Name:  "node",
	Usage: "Kontena Nodes",
	Flags: []cli.Flag{
		gridFlag,
	},
	Subcommands: []cli.Command{
		nodesListCommand,
	},
}
