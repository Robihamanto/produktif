package model

var (
	// ErrTodolistNotFound is used when DB queries does not find any user record
	ErrTodolistNotFound = AppError{
		ID:      10380001,
		Message: "Todolist not found",
	}
)

// Todolist represent todolist owe by user
type Todolist struct {
	Base
	User        User   `json:"user,omitempty"`
	UserID      uint   `json:"-" gorm:"not null"`
	Name        string `json:"name" gorm:"not null; size:255"`
	Description string `json:"description" gorm:"null; type:TEXT"`
	Tasks       []Task `json:"tasks"`
}

// TodolistDB represent all func to interact with todolist
type TodolistDB interface {
	View(id uint) (*Todolist, error)
	Create(*Todolist) (*Todolist, error)
	Update(*Todolist) (*Todolist, error)
	Delete(id uint) error
	List(id uint) ([]Todolist, error)
}
