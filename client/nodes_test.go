package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kontena/kontena-client-go/api"
	"github.com/kontena/kontena-client-go/client/test-data"
)

func TestNodeIDString(t *testing.T) {
	assert.Equal(t, "grid/name", NodeID{Grid: "grid", Name: "name"}.String(), "node ID string")
}

func testNodeIDParse(t *testing.T, expected NodeID, id string) {
	if nodeID, err := ParseNodeID(id); err != nil {
		t.Fatalf("parse %#v: %v", id, err)
	} else {
		assert.Equal(t, expected, nodeID, "node ID parse: %#v", id)
	}
}

func TestNodeIDParse(t *testing.T) {
	testNodeIDParse(t, NodeID{Grid: "grid", Name: "name"}, "grid/name")
}

func TestNodeGet(t *testing.T) {
	var test = makeTest()
	var testNode = test_data.Node

	test.mockGET("/v1/nodes/test/node1", "test-data/node.json")

	if node, err := test.client.Nodes.Get(NodeID{"test", "node1"}); err != nil {
		t.Fatalf("node get error: %v", err)
	} else {
		assert.Equal(t, testNode, node, "response node")
	}
}

func TestNodeGetToken(t *testing.T) {
	var test = makeTest()
	var testToken = api.NodeToken{
		ID:    "test/node",
		Token: "ZxeA2iQ1MT61oT808BG/ty6aKtSnsD4f1cUub+DHWTfKoCBLTVYuP/WrRyDvjZAWdHZ3jBf/mhjGMiWhJ4YpSg==",
	}

	test.mockGET("/v1/nodes/test/node/token", "test-data/node-token.json")

	if nodeToken, err := test.client.Nodes.GetToken(NodeID{"test", "node"}); err != nil {
		t.Fatalf("node get error: %v", err)
	} else {
		assert.Equal(t, nodeToken, testToken, "response node token")
	}
}

func testNodeCreate(t *testing.T, id NodeID, params api.NodePOST, mockRequest mockJSON) {
	var test = makeTest()

	test.mockPOST(t, "/v1/grids/"+id.Grid+"/nodes", func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "POST /v1/grids/"+id.Grid+"/nodes JSON")

		return api.Node{ID: id.String()}
	})

	if node, err := test.client.Nodes.Create(id.Grid, params); err != nil {
		t.Fatalf("nodes create error: %v", err)
	} else {
		assert.Equal(t, node.ID, id.String())
	}
}

func TestNodeCreate(t *testing.T) {
	var nodeParams = api.NodePOST{Name: "test-create"}
	var mockRequest = parseJSON(`{"name": "test-create"}`)

	testNodeCreate(t, NodeID{"test-grid", "test-create"}, nodeParams, mockRequest)
}

func TestNodeCreateParams(t *testing.T) {
	var nodeParams = api.NodePOST{
		Name:  "test-create-params",
		Token: "test-secret",
		Labels: &api.NodeLabels{
			"test==true",
		},
	}
	var mockRequest = parseJSON(`
		{
			"name": "test-create-params",
			"token": "test-secret",
			"labels": [ "test==true" ]
		}
	`)

	testNodeCreate(t, NodeID{"test-grid", "test-create-params"}, nodeParams, mockRequest)
}

func testNodeUpdate(t *testing.T, id NodeID, params api.NodePUT, mockRequest mockJSON) {
	var test = makeTest()

	test.mockPUT(t, "/v1/nodes/"+id.String(), func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "PUT /v1/nodes/"+id.String()+" JSON")

		return api.Node{ID: id.String()}
	})

	if node, err := test.client.Nodes.Update(id, params); err != nil {
		t.Fatalf("nodes update error: %v", err)
	} else {
		assert.Equal(t, node.ID, id.String())
	}
}

func TestNodeUpdateDefaults(t *testing.T) {
	var nodeParams = api.NodePUT{}
	var mockRequest = parseJSON(`{}`)

	testNodeUpdate(t, NodeID{"test-grid", "test-defaults"}, nodeParams, mockRequest)
}

func TestNodeUpdateLabelsClear(t *testing.T) {
	var nodeParams = api.NodePUT{
		Labels: &api.NodeLabels{},
	}
	var mockRequest = parseJSON(`{ "labels": [] }`)

	testNodeUpdate(t, NodeID{"test-grid", "test-labels-clear"}, nodeParams, mockRequest)
}

func TestNodeUpdateLabels(t *testing.T) {
	var nodeParams = api.NodePUT{
		Labels: &api.NodeLabels{"test==true"},
	}
	var mockRequest = parseJSON(`{ "labels": [ "test==true" ] }`)

	testNodeUpdate(t, NodeID{"test-grid", "test-labels"}, nodeParams, mockRequest)
}

func mockNodeTokenPUT(t *testing.T, id NodeID, mockRequest mockJSON) *test {
	var test = makeTest()

	test.mockPUT(t, "/v1/nodes/"+id.String()+"/token", func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "PUT /v1/nodes/"+id.String()+"/token JSON")

		return api.NodeToken{ID: id.String(), Token: "secret"}
	})

	return test
}
func mockNodeTokenDELETE(t *testing.T, id NodeID, mockRequest mockJSON) *test {
	var test = makeTest()

	test.mockDELETE(t, "/v1/nodes/"+id.String()+"/token", func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "DELETE /v1/nodes/"+id.String()+"/token JSON")

		return struct{}{}
	})

	return test
}

func TestNodeCreateTokenDefaults(t *testing.T) {
	var nodeID = NodeID{"test-grid", "test-node"}
	var params = api.NodeTokenParams{}
	var mockRequest = parseJSON(`{"reset_connection": false}`)

	var test = mockNodeTokenPUT(t, nodeID, mockRequest)

	if nodeToken, err := test.client.Nodes.CreateToken(nodeID, params); err != nil {
		t.Fatalf("nodes token update error: %v", err)
	} else {
		assert.Equal(t, nodeToken.ID, nodeID.String())
	}
}

func TestNodeCreateToken(t *testing.T) {
	var nodeID = NodeID{"test-grid", "test-node"}
	var params = api.NodeTokenParams{ResetConnection: true}
	var mockRequest = parseJSON(`{"reset_connection": true}`)

	var test = mockNodeTokenPUT(t, nodeID, mockRequest)

	if nodeToken, err := test.client.Nodes.CreateToken(nodeID, params); err != nil {
		t.Fatalf("nodes token update error: %v", err)
	} else {
		assert.Equal(t, nodeToken.ID, nodeID.String())
	}
}

func TestNodeUpdateToken(t *testing.T) {
	var nodeID = NodeID{"test-grid", "test-node"}
	var token = "secret 2"
	var params = api.NodeTokenParams{ResetConnection: true}
	var mockRequest = parseJSON(`{"reset_connection": true, "token": "secret 2"}`)

	var test = mockNodeTokenPUT(t, nodeID, mockRequest)

	if nodeToken, err := test.client.Nodes.UpdateToken(nodeID, token, params); err != nil {
		t.Fatalf("nodes token update error: %v", err)
	} else {
		assert.Equal(t, nodeToken.ID, nodeID.String())
	}
}

func TestNodeDeleteToken(t *testing.T) {
	var nodeID = NodeID{"test-grid", "test-node"}
	var params = api.NodeTokenParams{ResetConnection: true}
	var mockRequest = parseJSON(`{"reset_connection": true}`)

	var test = mockNodeTokenDELETE(t, nodeID, mockRequest)

	if err := test.client.Nodes.DeleteToken(nodeID, params); err != nil {
		t.Fatalf("nodes token delete error: %v", err)
	}
}
