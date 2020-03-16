package mysql

import (
	model "github.com/Robihamanto/produktif/internal"
	"github.com/jinzhu/gorm"
)

// TaskDB implementation of TaskDB interface
type TaskDB struct {
	cl *gorm.DB
}

// NewTaskDB create new instance of TaskDB
func NewTaskDB(db *gorm.DB) *TaskDB {
	return &TaskDB{db}
}

// View retrieve single task task id
func (d *TaskDB) View(id uint) (*model.Task, error) {
	var task model.Task
	if err := d.cl.Find(&task, id).Error; err != nil {
		return nil, model.ErrTaskNotFound
	} else if err != nil {
		return nil, err
	}
	return &task, nil
}

// Create new task for todolist
func (d TaskDB) Create(task *model.Task) (*model.Task, error) {
	err := d.cl.Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Update a task for todolist
func (d TaskDB) Update(task *model.Task) (*model.Task, error) {
	err := d.cl.Save(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Delete a task from todolist
func (d TaskDB) Delete(id uint) error {
	var t model.Task
	err := d.cl.Where("id = ?", id).Delete(&t).Error
	if err != nil {
		return err
	}
	return nil
}

// Unscope a task from todolist
func (d TaskDB) Unscope(id uint) error {
	var t model.Task
	err := d.cl.Unscoped().Where("id = ?", id).Delete(&t).Error
	if err != nil {
		return err
	}
	return nil
}

// List retrieve a bunch of task from todolist
func (d TaskDB) List(id uint) ([]model.Task, error) {
	var t []model.Task
	err := d.cl.Where("todolist_id = ?", id).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}
