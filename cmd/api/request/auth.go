package request

import "github.com/labstack/echo/v4"

// RegisterUser holds user registration request
type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

// ParseRegisterUser validates request body and parse it into RegisterUser
func ParseRegisterUser(c echo.Context) (*RegisterUser, error) {
	req := new(RegisterUser)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}
