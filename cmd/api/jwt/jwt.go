package jwt

import (
	"errors"
	"time"

	"github.com/Robihamanto/produktif/cmd/api/rbac"
)

// Service provides Json-Web-Token authentication implmentation
type Service struct {
	key      []byte
	duration time.Duration
	algo     string
	rbacSvc  *rbac.Service
}

var errInvalidToken = errors.New("Invalid token")

// TokenPayload represents token payload structure
type TokenPayload struct {
	UserID int
	Scope  string
}
