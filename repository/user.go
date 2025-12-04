package repository

import (
	"errors"
	"fmt"
	"main/database/model"

	"gorm.io/gorm"
)

var ErrUserExists = errors.New("user with given username or email already exists")
var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(input *model.User) (*model.User, error) {
	var registered model.User
	if err := r.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&registered).Error; err == nil {
		fmt.Println("User already exists:", registered.Username)
		return nil, ErrUserExists
	}

	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     model.RoleUser,
	}

	if err := user.SetPassword(input.Password); err != nil {
		return nil, err
	}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("username = ?", email).First(&user).Error; err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}
