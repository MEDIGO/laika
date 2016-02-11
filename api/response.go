package api

import (
	"net/http"

	"github.com/MEDIGO/feature-flag/store"
)

type Response struct {
	Status  int
	Payload interface{}
	Err     error
}

func (r Response) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return "no error"
}

func OK(payload interface{}) Response {
	return Response{http.StatusOK, payload, nil}
}

func Created(payload interface{}) Response {
	return Response{http.StatusCreated, payload, nil}
}

func NoContent() Response {
	return Response{http.StatusNoContent, nil, nil}
}

func BadRequest(err error) Response {
	return Response{http.StatusBadRequest, store.APIError{err.Error()}, err}
}

func NotFound(err error) Response {
	return Response{http.StatusNotFound, store.APIError{err.Error()}, err}
}

func InternalServerError(err error) Response {
	return Response{http.StatusInternalServerError, store.APIError{err.Error()}, err}
}
