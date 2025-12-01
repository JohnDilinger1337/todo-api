package service

import (
	"errors"

	"main/database/model"
	"main/repository"
)

type AuthService struct {
	UserRepo   *repository.UserRepository
	JWTService *JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *JWTService) *AuthService {
	return &AuthService{
		UserRepo:   userRepo,
		JWTService: jwtService,
	}
}

// Register a new user
func (s *AuthService) Register(username, email, password string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Email:    email,
		Role:     model.RoleUser,
		Password: password, // hashed in repository
	}

	createdUser, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// Login authenticates and returns a JWT token
func (s *AuthService) Login(username, password string) (string, error) {
	var user model.User
	if err := s.UserRepo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if !user.CheckPassword(password) {
		return "", errors.New("invalid password")
	}

	return s.JWTService.GenerateToken(user.ID, user.IsAdmin())
}
