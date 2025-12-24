package service

import (
	"errors"
	"strings"
	"time"

	"social-backend/internal/auth"
	"social-backend/internal/model"
	"social-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         *model.User `json:"user"`
}

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserRepo: repository.NewUserRepository(),
	}
}

// ==========================
// Register
// ==========================
func (s *AuthService) Register(email, password string) error {
	email = strings.ToLower(email)

	exist, _ := s.UserRepo.FindByEmail(email)
	if exist != nil {
		return errors.New("email already exists")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		Email:     email,
		Password:  string(hash),
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.UserRepo.Create(user)
}

// ==========================
// Login
// ==========================
func (s *AuthService) Login(email, password string) (*LoginResponse, error) {
	user, _ := s.UserRepo.FindByEmail(email)
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid email or password")
	}

	access, _ := auth.GenerateAccessToken(user.ID, user.Role)
	refresh, _ := auth.GenerateRefreshToken(user.ID)

	return &LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		User:         user,
	}, nil
}

// ==========================
// Refresh Token
// ==========================
func (s *AuthService) Refresh(refreshToken string) (string, error) {
	_, claims, err := auth.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	userID := uint(claims["user_id"].(float64))

	user, _ := s.UserRepo.FindByID(userID)
	if user == nil {
		return "", errors.New("user not found")
	}

	// issue new access token
	access, _ := auth.GenerateAccessToken(user.ID, user.Role)
	return access, nil
}
