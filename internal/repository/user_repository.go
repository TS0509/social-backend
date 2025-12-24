package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"social-backend/internal/database"
	"social-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (UserRepository) Create(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return database.DB.WithContext(ctx).Create(user).Error
}

func (UserRepository) FindByEmail(email string) (*model.User, error) {
	email = strings.ToLower(email)

	var user model.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := database.DB.Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (UserRepository) Update(user *model.User) error {
	return database.DB.Save(user).Error
}
