package auth

import (
	"log"

	model "github.com/Robihamanto/produktif/internal"
)

// New auth
func New(udb model.UserDB,
) *Service {
	return &Service{udb}
}

// Service represent authentication service
type Service struct {
	udb model.UserDB
}

// JWT is represent JWT interface
type JWT interface {
	GenerateToken(uint, model.AccessRole)
}

// RegisterUser is creating new user
func (s *Service) RegisterUser(username, password, email, fullname string) (*model.User, error) {

	// TODO: Check is user already registered
	user, err := s.udb.ViewByEmail(email)
	if err != nil && err != model.ErrUserNotFound {
		return nil, err
	}

	if user != nil {
		return nil, model.ErrUserAlreadyExist
	}

	u := &model.User{
		Username: username,
		Password: password,
		Email:    email,
		Fullname: fullname,
	}

	u, err = s.udb.Create(u)

	if err != nil {
		log.Print("Created user error: ", err)
		return nil, err
	}

	return u, nil
}
