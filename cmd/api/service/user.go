package service

import (
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	"github.com/Robihamanto/produktif/internal/user"
	echo "github.com/labstack/echo/v4"
)

// User represent user http service
type User struct {
	service *user.Service
}

// NewUser create service for user
func NewUser(
	service *user.Service,
	userRouter *echo.Group,
	jwtMw echo.MiddlewareFunc,
) {
	uhs := User{service}

	userRouter.GET("/:id", uhs.view, jwtMw)
	userRouter.GET("/me", uhs.viewMe, jwtMw)
}

func (u *User) view(c echo.Context) error {
	userID, err := request.ID(c)
	if err != nil {
		return err
	}
	result, err := u.service.View(uint(userID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (u *User) viewMe(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unable to fetch user",
		})
	}

	result, err := u.service.View(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
