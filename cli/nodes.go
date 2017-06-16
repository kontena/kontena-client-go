package cli

import (
	"fmt"
	"strings"

	"github.com/kontena/kontena-client-go/api"
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
	tbl := makeTable("Name", "Version", "Status", "Initial", "Labels")

	for _, node := range nodes {
		tbl.AddRow(node.Name,
			node.AgentVersion,
			nodeStatus(node),
			nodeInitial(node),
		)
	}

	tbl.Print()
}

type NodesCommand struct {
	*CLI
	Grid string
}

func (cmd NodesCommand) ListNodes() error {
	if nodes, err := cmd.client.Nodes.List(cmd.Grid); err != nil {
		return err
	} else {
		printNodes(nodes)
	}

	return nil
}
