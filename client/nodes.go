package client

import (
	"github.com/kontena/kontena-client-go/api"
)

type NodesAPI interface {
	List(grid string) ([]api.Node, error)
	Get(id api.NodeID) (api.Node, error)
	GetToken(id api.NodeID) (api.NodeToken, error)
	Create(grid string, params api.NodePOST) (api.Node, error)
	CreateToken(id api.NodeID, params api.NodeTokenParams) (api.NodeToken, error)
	Update(id api.NodeID, params api.NodePUT) (api.Node, error)
	UpdateToken(id api.NodeID, token string, params api.NodeTokenParams) (api.NodeToken, error)
	Delete(id api.NodeID) error
	DeleteToken(id api.NodeID, params api.NodeTokenParams) error
}

type nodesClient struct {
	client *Client
}

func (nodesClient nodesClient) List(grid string) ([]api.Node, error) {
	var nodes api.NodesGET

	return nodes.Nodes, nodesClient.client.get(request{ResponseBody: &nodes}, "/v1/grids", grid, "nodes")
}

func (nodesClient nodesClient) Get(id api.NodeID) (api.Node, error) {
	var node api.Node

	return node, nodesClient.client.get(request{ResponseBody: &node}, "/v1/nodes", id.String())
}

func (nodesClient nodesClient) GetToken(id api.NodeID) (api.NodeToken, error) {
	var nodeToken api.NodeToken

	return nodeToken, nodesClient.client.get(request{ResponseBody: &nodeToken}, "/v1/nodes", id.String(), "token")
}

func (nodesClient nodesClient) Create(grid string, params api.NodePOST) (api.Node, error) {
	var node api.Node

	return node, nodesClient.client.post(request{RequestBody: params, ResponseBody: &node}, "/v1/grids", grid, "nodes")
}

func (nodesClient nodesClient) CreateToken(id api.NodeID, params api.NodeTokenParams) (api.NodeToken, error) {
	var put = api.NodeTokenPUT{NodeTokenParams: params}
	var nodeToken api.NodeToken

	return nodeToken, nodesClient.client.put(request{RequestBody: put, ResponseBody: &nodeToken}, "/v1/nodes", id.String(), "token")
}

func (nodesClient nodesClient) Update(id api.NodeID, params api.NodePUT) (api.Node, error) {
	var node api.Node

	return node, nodesClient.client.put(request{RequestBody: params, ResponseBody: &node}, "/v1/nodes", id.String())
}

func (nodesClient nodesClient) UpdateToken(id api.NodeID, token string, params api.NodeTokenParams) (api.NodeToken, error) {
	var put = api.NodeTokenPUT{
		Token:           token,
		NodeTokenParams: params,
	}

	var nodeToken api.NodeToken

	return nodeToken, nodesClient.client.put(request{RequestBody: put, ResponseBody: &nodeToken}, "/v1/nodes", id.String(), "token")
}

func (nodesClient nodesClient) Delete(id api.NodeID) error {
	return nodesClient.client.delete(request{}, "/v1/nodes", id.String())
}

func (nodesClient nodesClient) DeleteToken(id api.NodeID, params api.NodeTokenParams) error {
	var requestBody = api.NodeTokenDELETE{
		NodeTokenParams: params,
	}

	return nodesClient.client.delete(request{RequestBody: requestBody}, "/v1/nodes", id.String(), "token")
}
