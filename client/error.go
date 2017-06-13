package client

import (
	"fmt"

	"github.com/kontena/terraform-provider-kontena/api"
)

type Error struct {
	HTTPStatus int
	API        api.Error
}

func (err Error) Error() string {
	return fmt.Sprintf("%v", err.API.Error)
}
