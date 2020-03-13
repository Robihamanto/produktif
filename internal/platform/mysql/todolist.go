package mysql

import (
	model "github.com/Robihamanto/produktif/internal"
	"github.com/jinzhu/gorm"
)

// TodolistDB implementation of TodolistDB interface
type TodolistDB struct {
	cl *gorm.DB
}

// NewTodolistDB create new instance of TodolistDB
func NewTodolistDB(db *gorm.DB) *TodolistDB {
	return &TodolistDB{db}
}

// View retrieve single todolist by user id
func (d *TodolistDB) View(id uint) (*model.Todolist, error) {
	var t model.Todolist
	if err := d.cl.Find(&t, id).Error; err == gorm.ErrRecordNotFound {
		return nil, model.ErrTodolistNotFound
	} else if err != nil {
		return nil, err
	}
	return &t, nil
}

// List is represent the list of todolist owe by user id
func (d *TodolistDB) List(id uint) ([]model.Todolist, error) {
	var t []model.Todolist
	err := d.cl.
		Where("user_id = ?", id).
		Find(&t).
		Error

	if err != nil {
		return nil, err
	}

	return t, nil
}

// Create is create new todolist from user
func (d *TodolistDB) Create(todolist *model.Todolist) (*model.Todolist, error) {
	var t model.Todolist
	err := d.cl.
		Create(todolist).
		Error

	if err != nil {
		return nil, err
	}

	return &t, nil
}

// Update is update todolist from user with new data
func (d *TodolistDB) Update(todolist *model.Todolist) ([]model.Todolist, error) {
	var t []model.Todolist
	err := d.cl.
		Create(todolist).
		Error

	if err != nil {
		return nil, err
	}

	return t, nil
}
