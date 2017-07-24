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
