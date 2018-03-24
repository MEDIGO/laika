package api

import (
	"github.com/labstack/echo"
)

func ListEnvironments(c echo.Context) error {
	return OK(c, getState(c).Environments)
}
