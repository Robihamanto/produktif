package model

import "golang.org/x/crypto/bcrypt"

var (
	// ErrUserAlreadyExist is used to inform that the user is already exist / duplicate
	ErrUserAlreadyExist = AppError{
		ID:      10280001,
		Message: "user already exist",
	}

	// ErrUserNotFound is used when DB queries does not find any user record
	ErrUserNotFound = AppError{
		ID:      10280002,
		Message: "User not found",
	}

	// ErrUserEmailAlreadyRegistered error
	ErrUserEmailAlreadyRegistered = AppError{
		ID:      10280003,
		Message: "User with that email already registered",
	}
)

// User represent model domain
type User struct {
	Base
	Password string `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Fullname string `json:"full_name"`
}

// HashPassword set the 'user' instance password which bcrypt-encrypted password
func (u *User) HashPassword(raw string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// VerifyPassword checks if the encrypted password match the plain-text password
func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// UserDB represent userdatabase (repository)
type UserDB interface {
	View(uint) (*User, error)
	ViewByEmail(string) (*User, error)
	Create(*User) (*User, error)
}
