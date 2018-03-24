package api

import (
	"github.com/MEDIGO/laika/models"
	"github.com/labstack/echo"
)

// RequestID returns the request ID from the current context.
func RequestID(c echo.Context) string {
	val := c.Get("request_id")
	if val == nil {
		return "unknown"
	}

	return val.(string)
}

func getState(c echo.Context) *models.State {
	state, _ := c.Get("state").(*models.State)
	return state
}
