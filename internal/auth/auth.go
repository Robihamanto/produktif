package auth

import (
	"log"

	model "github.com/Robihamanto/produktif/internal"
)

// Service represent authentication service
type Service struct {
	udb model.UserDB
	jwt JWT
}

// New auth
func New(
	udb model.UserDB,
	jwt JWT,
) *Service {
	return &Service{udb, jwt}
}

// JWT is represent JWT interface
type JWT interface {
	// GenerateToken(*model.User) (string, string, error)
	GenerateToken(uint, model.AccessRole) (string, string, error)
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

	var h model.Hashable = u
	err = h.HashPassword(password)
	if err != nil {
		return nil, err
	}

	u, err = s.udb.Create(u)

	if err != nil {
		log.Print("Created user error: ", err)
		return nil, err
	}

	return u, nil
}

// UserAuthentication wraps the result of Authenticate and AuthenticateOK
type UserAuthentication struct {
	User   *model.User `json:"user,omitempty"`
	Token  string      `json:"token"`
	Expiry string      `json:"expiry"`
}

// Authenticate tries to authenticate user from given username and password combination
func (s *Service) Authenticate(username, password string) (*UserAuthentication, error) {
	user, err := s.udb.ViewByUsername(username)
	if err != nil {
		return nil, err
	}

	err = user.VerifyPassword(password)
	if err != nil {
		return nil, model.ErrInvalidPassword
	}

	token, expiry, err := s.jwt.GenerateToken(user.ID, model.UserRole)
	if err != nil {
		return nil, err
	}

	return &UserAuthentication{user, token, expiry}, nil
}
