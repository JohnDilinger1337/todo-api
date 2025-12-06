package service

import (
	"main/config"
	"main/database/model"
	domainErr "main/domain/error"
	"main/repository"
)

type AuthService struct {
	UserRepo   *repository.UserRepository
	JWTService *JWTService
	Cfg        *config.Config
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

	if err := user.SetPassword(user.Password); err != nil {
		return nil, err
	}

	createdUser, err := s.UserRepo.CreateUser(user)

	return createdUser, err
}

// Login authenticates and returns a JWT token
func (s *AuthService) Login(username, password string) (string, error) {
	var user model.User
	if err := s.UserRepo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", domainErr.New(domainErr.ErrUserNotFoundCode)
	}

	if !user.CheckPassword(password) {
		return "", domainErr.New(domainErr.ErrInvalidPasswordCode)
	}

	token, err := s.JWTService.GenerateToken(user.ID, user.IsAdmin())
	if err != nil {
		return "", domainErr.New(domainErr.ErrTokenGenerationCode)
	}

	return token, nil

}
