package rbac

import (
	"net/http"

	"github.com/casbin/casbin"
	echo "github.com/labstack/echo/v4"
)

// Service is RBAC wrapper for API service
type Service struct {
	enforcer *casbin.Enforcer
}

// ErrInsufficientAccess error
var ErrInsufficientAccess = echo.NewHTTPError(http.StatusForbidden, "You are not allowed to access this resource")

// New creates new service instance
func New(e *casbin.Enforcer) *Service {
	return &Service{
		enforcer: e,
	}
}

// EnforceRole checks if the subject has permission to access
// corresponding resource (method and path)
func (s *Service) EnforceRole(subj, path, method string) error {
	if !s.enforcer.Enforce(subj, path, method) {
		return ErrInsufficientAccess
	}
	return nil
}

func (s *Service) checkRole(c echo.Context, role string) bool {
	r, ok := c.Get("user_role").(string)
	if !ok {
		return false
	}
	return r == role
}
