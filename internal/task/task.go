package task

import (
	"time"

	model "github.com/Robihamanto/produktif/internal"
)

// Service represent task service
type Service struct {
	taskRepo     model.TaskDB
	todolistRepo model.TodolistDB
}

// New creating new service for task
func New(
	taskRepo model.TaskDB,
	todolistRepo model.TodolistDB,
) *Service {
	return &Service{
		taskRepo,
		todolistRepo,
	}
}

// View retrieve single task from todolist
func (s *Service) View(id uint) (*model.Task, error) {
	result, err := s.taskRepo.View(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// List retrieve bunch of task from todolist
func (s *Service) List(id uint) ([]model.Task, error) {
	result, err := s.taskRepo.List(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Create hold task information for todolist
type Create struct {
	TodolistID  uint
	Title       string
	Description string
	DueDate     time.Time
	IsCompleted bool
}

// Create handle creation of task for todolist
func (s *Service) Create(param *Create) (*model.Task, error) {
	// Make sure todolist exist
	if _, err := s.todolistRepo.View(param.TodolistID); err != nil {
		return nil, err
	}

	task := &model.Task{
		TodolistID:  param.TodolistID,
		Title:       param.Title,
		Description: param.Description,
		DueDate:     param.DueDate,
		IsCompleted: param.IsCompleted,
	}

	result, err := s.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update hold task information for update task in todolist
type Update struct {
	TodolistID  *uint
	Title       *string
	Description *string
	DueDate     *time.Time
	IsCompleted *bool
}

// Update handle update information of task for todolist
func (s *Service) Update(id uint, param *Update) (*model.Task, error) {
	// get current task info
	task, err := s.taskRepo.View(id)

	if err != nil {
		return nil, err
	}

	if param.Title != nil {
		task.Title = *param.Title
	}

	if param.Description != nil {
		task.Description = *param.Description
	}

	if param.DueDate != nil {
		task.DueDate = *param.DueDate
	}

	if param.IsCompleted != nil {
		task.IsCompleted = *param.IsCompleted
	}

	result, err := s.taskRepo.Update(task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete is return error message if fail doing deletion
func (s *Service) Delete(id uint) error {
	err := s.taskRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
