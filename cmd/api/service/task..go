package service

import (
	"net/http"

	"github.com/Robihamanto/produktif/cmd/api/request"
	"github.com/Robihamanto/produktif/internal/task"
	"github.com/labstack/echo/v4"
)

// Task represent task service that hold all task made by user
type Task struct {
	service *task.Service
}

// NewTask creates new Task http service
func NewTask(
	service *task.Service,
	taskRouter *echo.Group,
	jwtMw echo.MiddlewareFunc,
) {
	ths := Task{service}

	taskRouter.GET("/todo/:id", ths.list, jwtMw)
	taskRouter.GET("/:id", ths.view, jwtMw)
	taskRouter.POST("", ths.create, jwtMw)
	taskRouter.PUT("/:id", ths.update, jwtMw)
	taskRouter.DELETE("/:id", ths.delete, jwtMw)
}

// GET /task/todo/:id
func (s *Task) list(c echo.Context) error {
	todolistID, err := request.ID(c)
	if err != nil {
		return err
	}

	result, err := s.service.List(uint(todolistID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// GET /task/:id
func (s *Task) view(c echo.Context) error {
	taskID, err := request.ID(c)
	if err != nil {
		return err
	}

	result, err := s.service.View(uint(taskID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// POST /task
func (s *Task) create(c echo.Context) error {
	req, err := request.ParseTask(c)
	if err != nil {
		return err
	}

	param := &task.Create{
		TodolistID:  req.TodolistID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		IsCompleted: req.IsCompleted,
	}

	result, err := s.service.Create(param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// PUT /task
func (s *Task) update(c echo.Context) error {
	taskID, err := request.ID(c)
	if err != nil {
		return err
	}

	req, err := request.ParseUpdateTask(c)
	if err != nil {
		return err
	}

	param := &task.Update{
		TodolistID:  req.TodolistID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		IsCompleted: req.IsCompleted,
	}

	result, err := s.service.Update(uint(taskID), param)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// DELETE /task/:id
func (s *Task) delete(c echo.Context) error {
	taskID, err := request.ID(c)
	if err != nil {
		return err
	}

	err = s.service.Delete(uint(taskID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete todolist success",
	})
}
