package client

import (
	"fmt"
	"strings"

	"github.com/kontena/kontena-client-go/api"
)

type NodeID struct {
	Grid string
	Name string
}

func ParseNodeID(id string) (nodeID NodeID, err error) {
	parts := strings.Split(id, "/")

	if len(parts) != 2 {
		return nodeID, fmt.Errorf("Invalid NodeID: %#v", id)
	}

	nodeID.Grid = parts[0]
	nodeID.Name = parts[1]

	return nodeID, nil
}

func (nodeID NodeID) String() string {
	return fmt.Sprintf("%s/%s", nodeID.Grid, nodeID.Name)
}

type NodesAPI interface {
	List(grid string) ([]api.Node, error)
	Get(id NodeID) (api.Node, error)
	Create(grid string, params api.NodePOST) (api.Node, error)
	Update(id NodeID, params api.NodePUT) (api.Node, error)
	Delete(id NodeID) error
}

type nodesClient struct {
	client *Client
}

func (nodesClient nodesClient) List(grid string) ([]api.Node, error) {
	var nodes api.NodesGET

	return nodes.Nodes, nodesClient.client.get(request{ResponseBody: &nodes}, "/v1/grids", grid, "nodes")
}

func (nodesClient nodesClient) Get(id NodeID) (api.Node, error) {
	var node api.Node

	return node, nodesClient.client.get(request{ResponseBody: &node}, "/v1/nodes", id.String())
}

func (nodesClient nodesClient) Create(grid string, params api.NodePOST) (api.Node, error) {
	var node api.Node

	return node, nodesClient.client.post(request{RequestBody: params, ResponseBody: &node}, "/v1/grids", grid, "nodes")
}

func (nodesClient nodesClient) Update(id NodeID, params api.NodePUT) (api.Node, error) {
	var node api.Node

	return node, nodesClient.client.put(request{RequestBody: params, ResponseBody: &node}, "/v1/nodes", id.String())
}

func (nodesClient nodesClient) Delete(id NodeID) error {
	return nodesClient.client.delete(request{}, "/v1/nodes", id.String())
}
