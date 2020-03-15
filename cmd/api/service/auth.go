package service

import (
	"log"
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	model "github.com/Robihamanto/produktif/internal"
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
	e.POST("/login", a.login)
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
		log.Print(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// login user is represent a way to user access the platform
func (a *Auth) login(c echo.Context) error {
	credentials, err := request.UserLogin(c)

	userAuth, err := a.svc.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		if err == model.ErrInvalidPassword {
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect username or password")
		}
	}

	return c.JSON(http.StatusOK, userAuth)
}
