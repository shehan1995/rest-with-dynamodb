package internal

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var (
	ErrNotFound            = &ErrResponse{HTTPStatusCode: 404, StatusText: "Item not found."}
	ErrBadRequest          = &ErrResponse{HTTPStatusCode: 400, StatusText: "Bad request"}
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: 500, StatusText: "Internal Server Error"}
)
