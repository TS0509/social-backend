package service

import (
	"errors"
	"strings"
	"time"

	"social-backend/internal/model"
	"social-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserRepo: repository.NewUserRepository(),
	}
}

func (s *AuthService) Register(email, password string) error {
	email = strings.ToLower(email)

	// Check exists
	existing, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("email already registered")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Email:     email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save
	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	email = strings.ToLower(email)

	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
