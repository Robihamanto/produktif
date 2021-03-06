package model

import "time"

// Base contains common fields for all tables
type Base struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Hashable serve interface to hash verivy hashed password
type Hashable interface {
	VerifyPassword(string) error
	HashPassword(string) error
}
