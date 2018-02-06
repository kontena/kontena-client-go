package client

import "github.com/kontena/kontena-client-go/api"

type ExternalRegistriesAPI interface {
	List(grid string) ([]api.ExternalRegistry, error)
	Get(id string) (api.ExternalRegistry, error)
	Create(grid string, params api.ExternalRegistryPOST) (api.ExternalRegistry, error)
	Delete(id string) error
}

type externalRegistriesClient struct {
	*Client
}

func (client externalRegistriesClient) List(grid string) ([]api.ExternalRegistry, error) {
	var index api.ExternalRegistryIndex

	return index.ExternalRegistries, client.get(request{ResponseBody: &index}, "/v1/grids", grid, "external_registries")
}

func (client externalRegistriesClient) Get(id string) (api.ExternalRegistry, error) {
	var externalRegistry api.ExternalRegistry

	return externalRegistry, client.get(request{ResponseBody: &externalRegistry}, "/v1/external_registries", id)
}

func (client externalRegistriesClient) Create(grid string, params api.ExternalRegistryPOST) (api.ExternalRegistry, error) {
	var externalRegistry api.ExternalRegistry

	return externalRegistry, client.post(request{RequestBody: params, ResponseBody: &externalRegistry}, "/v1/grids", grid, "external_registries")
}

func (client externalRegistriesClient) Delete(id string) error {
	return client.delete(request{}, "/v1/external_registries", id)
}
