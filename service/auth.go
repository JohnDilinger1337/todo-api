package service

import (
	"time"

	"main/config"
	"main/database/model"
	domainErr "main/domain/error"
	"main/repository"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	UserRepo   *repository.UserRepository
	JWTService *JWTService
	Cfg        *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *JWTService, cfg *config.Config) *AuthService {
	return &AuthService{
		UserRepo:   userRepo,
		JWTService: jwtService,
		Cfg:        cfg,
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
func (s *AuthService) Login(username, password string, c *gin.Context) error {
	var user model.User
	if err := s.UserRepo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return domainErr.New(domainErr.ErrUserNotFoundCode)
	}

	if !user.CheckPassword(password) {
		return domainErr.New(domainErr.ErrInvalidPasswordCode)
	}

	token, err := s.JWTService.GenerateToken(user.ID, user.IsAdmin())
	if err != nil {
		return domainErr.New(domainErr.ErrTokenGenerationCode)
	}

	duration, err := time.ParseDuration(s.Cfg.JWTExpiresAt)
	if err != nil {
		return domainErr.New(domainErr.ErrOther)
	}

	c.SetCookie(
		"token",                    // name
		token,                      // value (JWT token string)
		int(duration.Seconds()),    // max age in seconds (e.g., 1 hour)
		"/",                        // path
		"",                         // domain ("" means current domain)
		s.Cfg.GinMode == "release", // secure (only send over HTTPS)
		true,                       // httpOnly (not accessible via JS)
	)

	return nil

}
