package jwt

import (
	"errors"
	"net/http"
	"time"

	"github.com/Robihamanto/produktif/cmd/api/config"
	"github.com/Robihamanto/produktif/cmd/api/rbac"
	model "github.com/Robihamanto/produktif/internal"
	jwt "github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
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

// New create JWT instance
func New(c *config.JWT, rbacSvc *rbac.Service) *Service {
	return &Service{
		key:      []byte(c.Secret),
		duration: time.Duration(c.Duration) * time.Hour,
		algo:     c.SigningAlgorithm,
		rbacSvc:  rbacSvc,
	}
}

func (s *Service) authorize(c echo.Context, token *TokenPayload) error {
	path := c.Request().URL.Path
	method := c.Request().Method

	//token.Scope -> Role
	return s.rbacSvc.EnforceRole(token.Scope, path, method)
}

// AuthorizeTokenWithContext checks the token payload and bind it to the givne context
func (s *Service) AuthorizeTokenWithContext(token *TokenPayload, c echo.Context) error {
	err := s.authorize(c, token)
	if err != nil {
		return rbac.ErrInsufficientAccess
	}

	c.Set("user_id", token.UserID)
	c.Set("user_role", token.Scope)

	c.Set("is_authenticated", true)
	c.Set("is_authorize", true)

	return nil
}

// MWFunc is a middleware implementation for echo web framework
func (s *Service) MWFunc() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := s.ParseToken(c)
			if err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			err = s.AuthorizeTokenWithContext(token, c)
			if err != nil {
				return err
			}
			return next(c)
		}
	}
}

// ParseToken parses token from Authentication header
func (s *Service) ParseToken(c echo.Context) (*TokenPayload, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil, errors.New("token is not supplied")
	}

	// token without 'bearer' prefix
	return s.GetTokenPayload(token)
}

// GetTokenPayload generates TokenPayload from token string
func (s *Service) GetTokenPayload(token string) (*TokenPayload, error) {
	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(s.algo) != token.Method {
			return nil, errInvalidToken
		}
		return s.key, nil
	})

	if err != nil {
		return nil, err
	}

	if !payload.Valid {
		return nil, errInvalidToken
	}

	claims := payload.Claims.(jwt.MapClaims)

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errInvalidToken
	}

	scope := ""
	if claims["scope"] != nil {
		scope, ok = claims["scope"].(string)
		if !ok {
			return nil, errInvalidToken
		}
	}

	return &TokenPayload{
		UserID: int(userID),
		Scope:  scope,
	}, nil
}

// GenerateToken generates new JWT token with certain access type
func (s *Service) GenerateToken(userID uint, role model.AccessRole) (string, string, error) {
	duration := s.duration
	if role == model.AdminRole {
		duration = time.Hour * time.Duration(12)
	}

	expire := time.Now().Add(duration)

	token := jwt.NewWithClaims(jwt.GetSigningMethod(s.algo), jwt.MapClaims{
		"user_id": userID,
		"scope":   string(role),
		"exp":     expire.Unix(),
	})

	tokenS, err := token.SignedString(s.key)
	return tokenS, expire.Format(time.RFC3339), err
}
