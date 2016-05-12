package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Error represents an error message.
type Error struct {
	Message string `json:"message"`
}

// OK generates an HTTP 200 OK response with specified payload serialized as JSON.
func OK(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, payload)
}

// Created generates an HTTP 201 Created response with specified payload serialized as JSON.
func Created(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusCreated, payload)
}

// NoContent generates an empty HTTP 204 No Conent response.
func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// BadRequest generates an HTTP 400 Bad Request response with specified error serialized as JSON.
func BadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, Error{err.Error()})
}

// NotFound generates an HTTP 404 Not Found response with specified error serialized as JSON.
func NotFound(c echo.Context, err error) error {
	return c.JSON(http.StatusNotFound, Error{err.Error()})
}

// Conflict generates an HTTP 409 Conflict response with specified error serialized as JSON.
func Conflict(c echo.Context, err error) error {
	return c.JSON(http.StatusConflict, Error{err.Error()})
}

// InternalServerError generates an HTTP 500 Internal Server Error response with specified error serialized as JSON.
func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, Error{err.Error()})
}
