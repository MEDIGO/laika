package api

import "github.com/labstack/echo"

// RequestID returns the request ID from the current context.
func RequestID(c echo.Context) string {
	val := c.Get("request_id")
	if val == nil {
		return "unknown"
	}

	return val.(string)
}
