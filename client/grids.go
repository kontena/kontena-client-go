package client

import "github.com/kontena/kontena-client-go/api"

type GridsAPI interface {
	List() ([]api.Grid, error)
	Get(id string) (api.Grid, error)
	Create(params api.GridPOST) (api.Grid, error)
	Update(id string, params api.GridPUT) (api.Grid, error)
	Delete(id string) error
}

type gridsClient struct {
	client *Client
}

func (gridsClient gridsClient) List() ([]api.Grid, error) {
	var grids []api.Grid

	return grids, gridsClient.client.get(request{ResponseBody: &grids}, "/v1/grids")
}

func (gridsClient gridsClient) Get(id string) (api.Grid, error) {
	var grid api.Grid

	return grid, gridsClient.client.get(request{ResponseBody: &grid}, "/v1/grids", id)
}

func (gridsClient gridsClient) Create(params api.GridPOST) (api.Grid, error) {
	var grid api.Grid

	return grid, gridsClient.client.post(request{RequestBody: params, ResponseBody: &grid}, "/v1/grids")
}

func (gridsClient gridsClient) Update(id string, params api.GridPUT) (api.Grid, error) {
	var grid api.Grid

	return grid, gridsClient.client.put(request{RequestBody: params, ResponseBody: &grid}, "/v1/grids", id)
}

func (gridsClient gridsClient) Delete(id string) error {
	return gridsClient.client.delete(request{}, "/v1/grids", id)
}
