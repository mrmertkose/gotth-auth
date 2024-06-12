package repository

import (
	"github.com/mrmertkose/gotth-auth/model"
	"github.com/mrmertkose/gotth-auth/pkg/database"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(string) (*model.User, error)
	GetUserByID(uint) (*model.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func UserRepositoryNew() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (repo *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := repo.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
