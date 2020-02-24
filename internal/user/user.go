package user

import (
	model "github.com/Robihamanto/produktif/internal"
)

// Service represent user app service
type Service struct {
	udb model.UserDB
}

// New create new user app service
func New(udb model.UserDB) *Service {
	return &Service{udb}
}

// View is return single user by id
func (s *Service) View(id uint) (*model.User, error) {
	user, err := s.udb.View(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
