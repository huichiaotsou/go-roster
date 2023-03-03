package repository

import (
	"errors"

	"github.com/huichiaotsou/go-roster/pkg/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return result.Error
	}

	return nil
}

var ErrUserNotFound = errors.New("user not found")
