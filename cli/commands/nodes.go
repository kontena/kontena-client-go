package commands

import (
	"fmt"
	"strings"

	"github.com/kontena/kontena-client-go/api"
	"github.com/kontena/kontena-client-go/cli"
)

func nodeStatus(node api.Node) string {
	if !node.Connected {
		return "offline"
	}

	return "online"
}

func nodeInitial(node api.Node) string {
	if !node.InitialMember {
		return "-"
	}

	return fmt.Sprintf("%d / %d", node.NodeNumber, node.Grid.InitialSize)
}

func nodeLabels(node api.Node) string {
	return strings.Join(node.Labels, ",")
}

func printNodes(nodes []api.Node) {
	tbl := cli.Table("Name", "Version", "Status", "Initial", "Labels")

	for _, node := range nodes {
		tbl.AddRow(node.Name,
			node.AgentVersion,
			nodeStatus(node),
			nodeInitial(node),
		)
	}

	tbl.Print()
}

// NodesCommand represents the command type for a node within a grid.
type NodesCommand struct {
	*cli.CLI
	Grid string
}

// List prints a list of all nodes within the grid.
func (cmd NodesCommand) List() error {
	if nodes, err := cmd.Client.Nodes.List(cmd.Grid); err != nil {
		return err
	} else {
		printNodes(nodes)
	}

	return nil
}

// Show prints details about the given node within the grid.
func (cmd NodesCommand) Show(name string) error {
	if node, err := cmd.Client.Nodes.Get(api.NodeID{cmd.Grid, name}); err != nil {
		return err
	} else {
		return cli.Print(node)
	}
}
