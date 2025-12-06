package repository

import (
	"main/database/model"
	domainErr "main/domain/error"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {

	// Check duplicates
	var exists model.User
	if err := r.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&exists).Error; err == nil {
		return nil, &domainErr.DomainError{Code: domainErr.ErrUserExistsCode}
	}

	// Save to DB
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	user.Password = "" // never return hash
	return user, nil
}

// func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
// 	var user model.User
// 	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
// 		return nil, domainErr.New(domainErr.ErrUserNotFoundCode)
// 	}
// 	return &user, nil
// }

// func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
// 	var user model.User
// 	if err := r.DB.Where("username = ?", email).First(&user).Error; err != nil {
// 		return nil, domainErr.New(domainErr.ErrUserNotFoundCode)
// 	}
// 	return &user, nil
// }
