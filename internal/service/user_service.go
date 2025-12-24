package service

import (
	"errors"
	"social-backend/internal/model"
	"social-backend/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(),
	}
}

// GET profile
func (s *UserService) Profile(userID uint) (*model.User, error) {
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Update profile
func (s *UserService) Update(userID uint, avatar string) (*model.User, error) {
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Avatar = avatar
	err = s.UserRepo.Update(user)
	return user, err
}
