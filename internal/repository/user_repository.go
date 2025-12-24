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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user model.User
	err := database.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
