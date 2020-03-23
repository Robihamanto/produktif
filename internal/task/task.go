package task

import (
	"log"
	"time"

	model "github.com/Robihamanto/produktif/internal"
)

// Service represent task service
type Service struct {
	taskRepository model.TaskRepository
	todolistRepo   model.TodolistDB
}

// New creating new service for task
func New(
	taskRepo model.TaskRepository,
	todolistRepo model.TodolistDB,
) *Service {
	return &Service{
		taskRepo,
		todolistRepo,
	}
}

// View retrieve single task from todolist
func (s *Service) View(id uint) (*model.Task, error) {
	result, err := s.taskRepository.View(id)

	if err != nil {
		log.Print("Task service error: ", err)
		return nil, err
	}

	return result, nil
}

// List retrieve bunch of task from todolist
func (s *Service) List(id uint) ([]model.Task, error) {
	result, err := s.taskRepository.List(id)

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

	result, err := s.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Creates handle creation of task for todolist
func (s *Service) Creates(param []*Create) ([]*model.Task, error) {
	// Make sure todolist exist
	if _, err := s.todolistRepo.View(param[0].TodolistID); err != nil {
		return nil, err
	}

	var tasks []*model.Task

	for _, value := range param {
		task := &model.Task{
			TodolistID:  value.TodolistID,
			Title:       value.Title,
			Description: value.Description,
			DueDate:     value.DueDate,
			IsCompleted: value.IsCompleted,
		}
		tasks = append(tasks, task)
	}

	result, err := s.taskRepository.Creates(tasks)
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
	task, err := s.taskRepository.View(id)

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

	result, err := s.taskRepository.Update(task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete is return error message if fail doing deletion
func (s *Service) Delete(id uint) error {
	err := s.taskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Unscope is return error message if fail doing deletion
func (s *Service) Unscope(id uint) error {
	err := s.taskRepository.Unscope(id)
	if err != nil {
		return err
	}
	return nil
}
