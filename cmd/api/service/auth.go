package service

import (
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	"github.com/Robihamanto/produktif/internal/auth"
	"github.com/labstack/echo/v4"
)

// Auth represent auth http service
type Auth struct {
	svc *auth.Service
}

// NewAuth create new authentication http service
func NewAuth(svc *auth.Service, e *echo.Echo) {
	a := Auth{svc}

	e.POST("/register", a.register)
}

// register is creating new user
func (a *Auth) register(c echo.Context) error {
	req, err := request.ParseRegisterUser(c)
	if err != nil {
		return err
	}

	user, err := a.svc.RegisterUser(
		req.Username,
		req.Password,
		req.Email,
		req.Fullname,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
