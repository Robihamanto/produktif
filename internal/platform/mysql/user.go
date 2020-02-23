package mysql

import (
	"github.com/jinzhu/gorm"
	model "gitlab.com/comeapp/comeapp-backend/internal"
)

// UserDB represent client for the user table
func UserDB struct {
	cl *gorm.DB
}

// View return single user by ID
func (u *UserDB) View(id uint) (*model.User, error) {
	var user model.User
	if err := u.cl.Find(&user,id).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// View return single user by ID
func (u *UserDB) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.cl.Where("email = ?", email).Find(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) Delete(user *model.User) error {
	if err := u.cl.Delete(user).Error; err != nil {
		return err
	}
	return nil
}