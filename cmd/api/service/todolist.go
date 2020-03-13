package service

import (
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	model "github.com/Robihamanto/produktif/internal"
	"github.com/Robihamanto/produktif/internal/todolist"
	"github.com/labstack/echo/v4"
)

type todolistHTTPService struct {
	svc *todolist.Service
}

// NewTodolist creates new Todolist http service
// svc : Service
// tr : todolistRouter
// jwtMw : jwtMiddleware
func NewTodolist(
	svc *todolist.Service,
	tr *echo.Group,
	jwtMw echo.MiddlewareFunc,
) {
	ths := todolistHTTPService{svc}

	tr.GET("", ths.list, jwtMw)
	tr.POST("", ths.create)
	tr.PUT("/:id", ths.create)
}

func (s *todolistHTTPService) list(c echo.Context) error {
	userID, err := request.ID(c)
	if err != nil {
		return err
	}

	var result []model.Todolist
	result, err = s.svc.List(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (s *todolistHTTPService) create(c echo.Context) error {
	req, err := request.ParseTodolist(c)
	if err != nil {
		return err
	}

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return model.ErrCastingFailure
	}

	param := &todolist.Create{
		UserID:      uint(userID),
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := s.svc.Create(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (s *todolistHTTPService) update(c echo.Context) error {
	todolistID, err := request.ID(c)
	if err != nil {
		return err
	}

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return model.ErrCastingFailure
	}

	req, err := request.ParseUpdateTodolist(c)
	if err != nil {
		return err
	}

	t := &todolist.Update{
		Name:        req.Name,
		Description: req.Description,
	}

	res, err := s.svc.Update(uint(todolistID), uint(userID), t)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
