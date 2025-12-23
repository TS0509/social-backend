package service

import (
	"errors"
	"social-backend/internal/model"
	"social-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Register(email, password string) error {
	// Hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		Email:    email,
		Password: string(hashed),
	}

	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
