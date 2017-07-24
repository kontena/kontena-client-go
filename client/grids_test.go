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

func TestGridCreateOptions(t *testing.T) {
	var test = makeTest()
	var gridParams = api.GridPOST{
		Name:        "test-options",
		InitialSize: 3,
		Token:       "secret 2",
		Subnet:      "10.8.0.0/16",
		Supernet:    "10.0.0.0/8",
		TrustedSubnets: &api.GridTrustedSubnets{
			"192.168.66.0/24",
		},
		DefaultAffinity: &api.GridDefaultAffinity{"test==true"},
		Stats: &api.GridStats{
			Statsd: &api.GridStatsStatsd{
				Server: "127.0.0.1",
				Port:   8125,
			},
		},
		Logs: &api.GridLogs{
			Forwarder: "test",
			Opts: api.GridLogsOpts{
				"server": "localhost",
			},
		},
	}
	var mockRequest = parseJSON(`
    {
      "name":             "test-options",
      "initial_size":     3,
      "token":            "secret 2",
      "subnet":           "10.8.0.0/16",
      "supernet":         "10.0.0.0/8",
      "trusted_subnets":  ["192.168.66.0/24"],
      "default_affinity": ["test==true"],
      "logs": {
        "forwarder": "test",
        "opts": { "server": "localhost" }
      },
      "stats": {
        "statsd": {
          "server": "127.0.0.1",
          "port": 8125
        }
      }
    }
  `)

	test.mockPOST(t, "/v1/grids", func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "POST /v1/grids JSON")

		return api.Grid{ID: "test"}
	})

	if grid, err := test.client.Grids.Create(gridParams); err != nil {
		t.Fatalf("grids create error: %v", err)
	} else {
		assert.Equal(t, grid.ID, "test")
	}
}
