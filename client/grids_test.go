package client

import (
	"testing"

	"github.com/kontena/kontena-client-go/api"
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

func TestGridCreate(t *testing.T) {
	var test = makeTest()
	var gridParams = api.GridPOST{
		Name:        "test",
		InitialSize: 1,
	}
	var mockRequest = parseJSON(`
    {
      "name":         "test",
      "initial_size": 1
    }
  `)

	test.mockPOST(t, "/v1/grids", func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "POST /v1/grids JSON")

		return api.Grid{
			ID:          "test",
			Name:        "test",
			InitialSize: 1,
			Token:       "secret",
		}
	})

	if grid, err := test.client.Grids.Create(gridParams); err != nil {
		t.Fatalf("grids create error: %v", err)
	} else {
		assert.Equal(t, grid.ID, "test")
		assert.Equal(t, grid.Token, "secret")
	}
}
