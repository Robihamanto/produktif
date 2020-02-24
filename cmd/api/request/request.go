package request

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ID return id url parameyer
// In case of conversion error to int, StatusBadRequest will be returned ass err
func ID(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest)
	}
	return id, nil
}
