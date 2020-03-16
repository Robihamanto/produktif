package request

import (
	"log"

	"github.com/labstack/echo/v4"
)

// CreateTodolist holds information of todolist to be created
type CreateTodolist struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=65535"`
}

// UpdateTodolist holds information of todolist to be created
type UpdateTodolist struct {
	Name        *string `json:"name" validate:"required,max=255"`
	Description *string `json:"description" validate:"required,max=65535"`
}

// ParseTodolist parses http request and save the information to
// Todolist Struct
func ParseTodolist(c echo.Context) (*CreateTodolist, error) {
	p := new(CreateTodolist)
	if err := c.Bind(p); err != nil {
		log.Print("Error binding when create: ", err)
		return nil, err
	}
	return p, nil
}

// ParseUpdateTodolist parse http request body and map it into struct
func ParseUpdateTodolist(c echo.Context) (*UpdateTodolist, error) {
	u := new(UpdateTodolist)
	if err := c.Bind(u); err != nil {
		log.Print("Error binding when update: ", err)
		return nil, err
	}
	return u, nil
}
