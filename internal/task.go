package model

import "time"

var (
	// ErrTaskNotFound is used when DB queries does not find any user record
	ErrTaskNotFound = AppError{
		ID:      10480001,
		Message: "Task not found",
	}
)

// Task task todolist owe by todolist
type Task struct {
	Base
	Todolist    Todolist  `json:"todolist,omitempty"`
	TodolistID  uint      `json:"-" gorm:"not null"`
	Title       string    `json:"title" gorm:"null"`
	Description string    `json:"description" gorm:"not null; type:TEXT"`
	DueDate     time.Time `json:"due_date" gorm:"not null;default:'1971-01-01 00:00:00'"`
	IsCompleted bool      `json:"is_completed" gorm:"not null"`
}

// TaskDB represent all function to interact with Task database
type TaskDB interface {
	View(id uint) (*Task, error)
	Create(*Task) (*Task, error)
	Update(*Task) (*Task, error)
	Delete(id uint) error
	Unscope(id uint) error
	List(id uint) ([]Task, error)
}
