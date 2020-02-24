package model

// AccessRole is represent access the role type
type AccessRole string

const (
	// AdminRole has control to almost all resources
	AdminRole AccessRole = "admin"
	// UserRole has control to almost all resources
	UserRole AccessRole = "user"
)
