package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
