package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status_text"`
	Message        string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrResponse{HTTPStatusCode: 405, StatusText: "Method Not Allowed", Message: "HTTP Method not allowed"}
	ErrNotFound         = &ErrResponse{HTTPStatusCode: 404, StatusText: "Not Found", Message: "Resource not found"}
	ErrBadRequest       = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad Request", Message: "Bad request"}
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Bad Request",
		Message:        err.Error(),
	}
}

func ServerErrorRender(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error",
		Message:        err.Error(),
	}
}

func NotFoundRender(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		StatusText:     "Not Found",
		Message:        err.Error(),
	}
}

func MethodNotAllowedRender(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 405,
		StatusText:     "Method Not Allowed",
		Message:        err.Error(),
	}
}
