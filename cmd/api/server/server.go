package server

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// New instance of Server
func New() *echo.Echo {
	e := echo.New()

	e.GET("/", healthCheck)

	return e
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status": "OK",
	})
}
