package model

var (
	// ErrInvalidPassword is used to inform that given password doesn't match
	// the hash
	ErrInvalidPassword = AppError{
		ID:      10040001,
		Message: "invalid password",
	}

	// ErrInvalidResetPasswordToken error
	ErrInvalidResetPasswordToken = AppError{
		ID:      10040002,
		Message: "Invalid email or token",
	}
)

// AccessRole is represent access the role type
type AccessRole string

const (
	// AdminRole has control to almost all resources
	AdminRole AccessRole = "admin"
	// UserRole has control to almost user resources
	UserRole AccessRole = "user"
)
