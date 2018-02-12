package client

import "github.com/kontena/kontena-client-go/api"

type ExternalRegistriesAPI interface {
	List(grid string) ([]api.ExternalRegistry, error)
	Get(id api.ExternalRegistryID) (api.ExternalRegistry, error)
	Create(grid string, params api.ExternalRegistryPOST) (api.ExternalRegistry, error)
	Delete(id api.ExternalRegistryID) error
}

type externalRegistriesClient struct {
	*Client
}

func (client externalRegistriesClient) List(grid string) ([]api.ExternalRegistry, error) {
	var index api.ExternalRegistryIndex

	return index.ExternalRegistries, client.get(request{ResponseBody: &index}, "/v1/grids", grid, "external_registries")
}

func (client externalRegistriesClient) Get(id api.ExternalRegistryID) (api.ExternalRegistry, error) {
	var externalRegistry api.ExternalRegistry

	return externalRegistry, client.get(request{ResponseBody: &externalRegistry}, "/v1/external_registries", id.Grid, id.Name)
}

func (client externalRegistriesClient) Create(grid string, params api.ExternalRegistryPOST) (api.ExternalRegistry, error) {
	var externalRegistry api.ExternalRegistry

	return externalRegistry, client.post(request{RequestBody: params, ResponseBody: &externalRegistry}, "/v1/grids", grid, "external_registries")
}

func (client externalRegistriesClient) Delete(id api.ExternalRegistryID) error {
	return client.delete(request{}, "/v1/external_registries", id.Grid, id.Name)
}
