package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
