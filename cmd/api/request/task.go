package request

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// CreateTask holds information of todolist to be created
type CreateTask struct {
	TodolistID  uint      `json:"todolist_id" validate:"required"`
	Title       string    `json:"title" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=65535"`
	DueDate     time.Time `json:"due_date" validate:"default:'1971-01-01 00:00:00'"`
	IsCompleted bool      `json:"is_completed" validate:"omitempty"`
}

// UpdateTask holds information of todolist to be created
type UpdateTask struct {
	TodolistID  *uint      `json:"todolist_id" validate:"required"`
	Title       *string    `json:"title" validate:"required,max=255"`
	Description *string    `json:"description" validate:"required,max=65535"`
	DueDate     *time.Time `json:"due_date" validate:"default:'1971-01-01 00:00:00'"`
	IsCompleted *bool      `json:"is_completed" validate:"omitempty"`
}

// ParseTask parses http request and save the information to
// Task Struct
func ParseTask(c echo.Context) (*CreateTask, error) {
	p := new(CreateTask)
	if err := c.Bind(p); err != nil {
		log.Print("Error binding when create: ", err)
		return nil, err
	}
	return p, nil
}

// ParseUpdateTask parse http request body and map it into struct
func ParseUpdateTask(c echo.Context) (*UpdateTask, error) {
	u := new(UpdateTask)
	if err := c.Bind(u); err != nil {
		log.Print("Error binding when update: ", err)
		return nil, err
	}
	return u, nil
}
