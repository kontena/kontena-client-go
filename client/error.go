package client

import (
	"fmt"
	"net/http"

	"github.com/kontena/kontena-client-go/api"
)

type Error struct {
	httpRequest  *http.Request
	httpResponse *http.Response
	apiError     api.Error
}

func (err Error) HTTPStatus() int {
	return err.httpResponse.StatusCode
}
func (err Error) APIError() api.Error {
	return err.apiError
}

func (err Error) Error() string {
	return fmt.Sprintf("%v %v => HTTP %v: %v",
		err.httpRequest.Method, err.httpRequest.URL,
		err.httpResponse.Status,
		err.apiError.Error,
	)
}

type NotFoundError Error
type ForbiddenError Error

func (err NotFoundError) Error() string {
	return Error(err).Error()
}
func (err ForbiddenError) Error() string {
	return Error(err).Error()
}
