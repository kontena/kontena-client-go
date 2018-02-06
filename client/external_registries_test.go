package client

import (
	"testing"

	"github.com/kontena/kontena-client-go/api"
	"github.com/stretchr/testify/assert"
)

func TestExternalRegistryList(t *testing.T) {
	var test = makeTest()

	test.mockGETFile("/v1/grids/test/external_registries", "test-data/external-registries.json")

	if externalRegistries, err := test.client.ExternalRegistries.List("test"); err != nil {
		t.Fatalf("grids list error: %v", err)
	} else {
		assert.Equal(t, externalRegistries, []api.ExternalRegistry{
			api.ExternalRegistry{
				ID:       "test/images.kontena.io",
				Name:     "images.kontena.io",
				URL:      "https://images.kontena.io",
				Username: "test",
			},
		})
	}
}

func TestExternalRegistryGet(t *testing.T) {
	var test = makeTest()

	test.mockGET("/v1/external_registries/test/images.kontena.io", `{"id":"test/images.kontena.io","name":"images.kontena.io","url":"https://images.kontena.io","username":"test","email":null}`)

	if externalRegistry, err := test.client.ExternalRegistries.Get("test/images.kontena.io"); err != nil {
		t.Fatalf("external-registry get error: %v", err)
	} else {
		assert.Equal(t, externalRegistry, api.ExternalRegistry{
			ID:       "test/images.kontena.io",
			Name:     "images.kontena.io",
			URL:      "https://images.kontena.io",
			Username: "test",
		}, "response node")
	}
}

func TestExternalRegistryCreate(t *testing.T) {
	var test = makeTest()
	var params = api.ExternalRegistryPOST{
		URL:      "https://images.kontena.io",
		Username: "test",
		Password: "password",
	}
	var mockRequest = parseJSON(`
    {"username":"test","password":"password","url":"https://images.kontena.io"}
  `)

	test.mockPOST(t, "/v1/grids/test/external_registries", func(request mockJSON) interface{} {
		assert.Equal(t, mockRequest, request, "POST /v1/grids/../external_registries JSON")

		return api.ExternalRegistry{
			ID:       "test/images.kontena.io",
			Name:     "images.kontena.io",
			URL:      "https://images.kontena.io",
			Username: "test",
		}
	})

	if externalRegistry, err := test.client.ExternalRegistries.Create("test", params); err != nil {
		t.Fatalf("external-registries create error: %v", err)
	} else {
		assert.Equal(t, externalRegistry.ID, "test/images.kontena.io")
	}
}
