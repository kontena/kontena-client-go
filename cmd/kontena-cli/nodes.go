package main

import (
	"fmt"
	"strings"

	"github.com/kontena/kontena-client-go/api"
	"github.com/rodaine/table"
	"github.com/urfave/cli"
)

func nodeStatus(node api.Node) string {
	if node.Connected {
		return "online"
	} else {
		return "offline"
	}
}

func nodeInitial(node api.Node) string {
	if node.InitialMember {
		return fmt.Sprintf("%d / %d", node.NodeNumber, node.Grid.InitialSize)
	} else {
		return "-"
	}
}

func nodeLabels(node api.Node) string {
	return strings.Join(node.Labels, ",")
}

func printNodes(nodes []api.Node) {
	tbl := table.New("Name", "Version", "Status", "Initial", "Labels")

	for _, node := range nodes {
		tbl.AddRow(node.Name,
			node.AgentVersion,
			nodeStatus(node),
			nodeInitial(node),
		)
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
