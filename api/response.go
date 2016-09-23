package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// Error represents an error message response.
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

// BadRequest generates an HTTP 400 Bad Request response with specified error message serialized as JSON.
func BadRequest(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, Error{msg})
}

// NotFound generates an HTTP 404 Not Found response with a generic error message serialized as JSON.
func NotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, Error{"Resource not found"})
}

// Conflict generates an HTTP 409 Conflict response with specified error message serialized as JSON.
func Conflict(c echo.Context, msg string) error {
	return c.JSON(http.StatusConflict, Error{msg})
}

// Invalid generates an HTTP 422 Unprocessable Entity response with specified error message serialized as JSON.
func Invalid(c echo.Context, msg string) error {
	// this status is not in the net/http package
	return c.JSON(422, Error{msg})
}

// InternalServerError generates an HTTP 500 Internal Server Error response with a generic error message
// serialized as JSON, while the provided error is logged with ERROR level.
func InternalServerError(c echo.Context, err error) error {
	log.WithFields(log.Fields{
		"request_id": RequestID(c),
	}).Error(err)

	return c.JSON(http.StatusInternalServerError, Error{"Oops! Something went wrong"})
}
