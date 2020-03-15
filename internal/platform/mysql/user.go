package mysql

import (
	model "github.com/Robihamanto/produktif/internal"
	"github.com/jinzhu/gorm"
)

// UserDB represent client for the user table
type UserDB struct {
	cl *gorm.DB
}

// NewUserDB returning a new UserDB instance
func NewUserDB(c *gorm.DB) *UserDB {
	return &UserDB{c}
}

// View return single user by ID
func (u *UserDB) View(id uint) (*model.User, error) {
	var user model.User
	if err := u.cl.Find(&user, id).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create update user on database
func (u *UserDB) Create(user *model.User) (*model.User, error) {
	if err := u.cl.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ViewByEmail return single user by email
func (u *UserDB) ViewByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.cl.
		Where("email = ?", email).
		Find(&user).
		Error; err == gorm.ErrRecordNotFound {
		return nil, model.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// ViewByUsername return single user by username
func (u *UserDB) ViewByUsername(username string) (*model.User, error) {
	var user model.User
	if err := u.cl.
		Where("username = ?", username).
		Find(&user).
		Error; err == gorm.ErrRecordNotFound {
		return nil, model.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete return single user by user data
func (u *UserDB) Delete(user *model.User) error {
	if err := u.cl.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
