package todolist

import (
	model "github.com/Robihamanto/produktif/internal"
)

// Service represent user app service
type Service struct {
	todolistRepo model.TodolistDB
	userRepo     model.UserDB
}

// New create new user app service
func New(
	todolistRepo model.TodolistDB,
	userRepo model.UserDB,
) *Service {
	return &Service{
		todolistRepo,
		userRepo,
	}
}

// View is return single user by id
func (s *Service) View(id uint) (*model.Todolist, error) {
	user, err := s.todolistRepo.View(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// List is return list of todolist by id
func (s *Service) List(id uint) ([]model.Todolist, error) {
	todolists, err := s.todolistRepo.List(id)
	if err != nil {
		return nil, err
	}
	return todolists, nil
}

// Create struct for todolist
type Create struct {
	UserID      uint
	Name        string
	Description string
}

// Create new Todolist instance and save it into DB
func (s *Service) Create(param *Create) (*model.Todolist, error) {
	// Make sure user exist
	_, err := s.userRepo.View(param.UserID)
	if err != nil {
		return nil, err
	}

	todolist := &model.Todolist{
		UserID:       param.UserID,
		Name:         param.Name,
		Desctription: param.Description,
	}

	return todolist, err
}

// Update struct hold data for update todolist
type Update struct {
	UserID      *uint
	Name        *string
	Description *string
}

// Update new Todolist instance and save it into DB
func (s *Service) Update(id, userID uint, param *Update) (*model.Todolist, error) {
	// Get current todolist
	todolist, err := s.todolistRepo.View(id)
	if err != nil {
		return nil, err
	}

	if param.Name != nil {
		todolist.Name = *param.Name
	}

	if param.Description != nil {
		todolist.Desctription = *param.Description
	}

	return todolist, err
}
