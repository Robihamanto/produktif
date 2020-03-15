package service

import (
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	model "github.com/Robihamanto/produktif/internal"
	"github.com/Robihamanto/produktif/internal/todolist"
	"github.com/labstack/echo/v4"
)

// Todolist represent todolist that hold all task made by user
type Todolist struct {
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
	ths := Todolist{svc}

	tr.GET("", ths.list, jwtMw)
	tr.GET("/:id", ths.view, jwtMw)
	tr.POST("", ths.create, jwtMw)
	tr.PUT("/:id", ths.update, jwtMw)
	tr.DELETE("/:id", ths.delete, jwtMw)
}

// GET /todolist
func (s *Todolist) list(c echo.Context) error {

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return model.ErrCastingFailure
	}

	//var result []model.Todolist
	result, err := s.svc.List(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// GET /todolist/:id
func (s *Todolist) view(c echo.Context) error {
	todolistID, err := request.ID(c)
	if err != nil {
		return err
	}

	result, err := s.svc.View(uint(todolistID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// POST /todolist
func (s *Todolist) create(c echo.Context) error {
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

// PUT /todolist
func (s *Todolist) update(c echo.Context) error {
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

// DELETE /todolist/:id
func (s *Todolist) delete(c echo.Context) error {
	todolistID, err := request.ID(c)
	if err != nil {
		return err
	}

	err = s.svc.Delete(uint(todolistID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete todolist success",
	})
}
