package client

import (
	"fmt"
	"net/http"

	"github.com/kontena/terraform-provider-kontena/api"
)

type Error struct {
	httpRequest  *http.Request
	httpResponse *http.Response
	API          api.Error
}

func (err Error) Error() string {
	return fmt.Sprintf("%v %v => HTTP %v %v: %v",
		err.httpRequest.Method, err.httpRequest.URL,
		err.httpResponse.StatusCode, err.httpResponse.Status,
		err.API.Error,
	)
}

type NotFoundError Error

func (err NotFoundError) Error() string {
	return Error(err).Error()
}
