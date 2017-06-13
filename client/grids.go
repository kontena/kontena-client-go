package client

import (
	"github.com/dghubble/sling"

	"github.com/kontena/terraform-provider-kontena/api"
)

type GridsAPI interface {
	List() ([]api.Grid, error)
	Get(id string) (api.Grid, error)
	Create(params api.GridPOST) (api.Grid, error)
	Update(id string, params api.GridPUT) (api.Grid, error)
	Delete(id string) error
}

type gridsClient struct {
	sling *sling.Sling
}

func (gridsClient gridsClient) List() ([]api.Grid, error) {
	var grids []api.Grid

	return grids, do(gridsClient.sling.New().Get("/v1/grids"), &grids)
}

func (gridsClient gridsClient) Get(id string) (api.Grid, error) {
	var grid api.Grid

	return grid, do(gridsClient.sling.New().Path("/v1/grids/").Get(id), &grid)
}

func (gridsClient gridsClient) Create(params api.GridPOST) (api.Grid, error) {
	var grid api.Grid

	return grid, do(gridsClient.sling.New().Post("/v1/grids").BodyJSON(params), &grid)
}

func (gridsClient gridsClient) Update(id string, params api.GridPUT) (api.Grid, error) {
	var grid api.Grid

	return grid, do(gridsClient.sling.New().Path("/v1/grids/").Put(id).BodyJSON(params), &grid)
}

func (gridsClient gridsClient) Delete(id string) error {
	return do(gridsClient.sling.New().Path("/v1/grids/").Delete(id), nil)
}
