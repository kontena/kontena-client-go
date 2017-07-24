package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridsList(t *testing.T) {
	var test = makeTest()

	test.mockGET("/v1/grids", "test-data/grids.json")

	if grids, err := test.client.Grids.List(); err != nil {
		t.Fatalf("grids list error: %v", err)
	} else {
		assert.Equal(t, len(grids), 1, "array len")
		assert.Equal(t, grids[0].ID, "test", "grid id")
	}
}
