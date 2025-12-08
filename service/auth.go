package service

import (
	"main/database/model"
	domainErr "main/domain/error"
	"main/dto"
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
func (s *AuthService) Register(input *dto.RegisterInput) (*model.User, error) {
	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     model.RoleUser,
		Password: input.Password, // hashed in repository
	}

	if err := user.SetPassword(user.Password); err != nil {
		return nil, err
	}

	createdUser, err := s.UserRepo.CreateUser(user)

	return createdUser, err
}

// Login authenticates and returns a JWT token
func (s *AuthService) Login(input *dto.LoginInput) (string, error) {
	var user model.User
	if err := s.UserRepo.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return "", domainErr.New(domainErr.ErrUserNotFoundCode)
	}

	if !user.CheckPassword(input.Password) {
		return "", domainErr.New(domainErr.ErrInvalidPasswordCode)
	}

	token, err := s.JWTService.GenerateToken(user.ID, user.IsAdmin())
	if err != nil {
		return "", domainErr.New(domainErr.ErrTokenGenerationCode)
	}

	return token, nil

}
