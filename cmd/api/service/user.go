package service

import (
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	"github.com/Robihamanto/produktif/internal/user"
	echo "github.com/labstack/echo/v4"
)

// userHTTPService represent user http service
type userHTTPService struct {
	svc *user.Service
}

// NewUser create service for user
func NewUser(
	svc *user.Service,
	ur *echo.Group,
) {
	uhs := userHTTPService{svc}

	ur.GET("/:id", uhs.view)
}

func (u *userHTTPService) view(c echo.Context) error {
	userID, err := request.ID(c)
	if err != nil {
		return err
	}
	result, err := u.svc.View(uint(userID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}
