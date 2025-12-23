package repository

import (
	"social-backend/internal/database"
	"social-backend/internal/model"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
