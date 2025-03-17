package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/models"
)

type UserRepository interface {
	GetById(id uuid.UUID) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	// Login(email string, password string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetById(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, "email = ?", email).Error; err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	fmt.Println("userss", user)
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {

	if database.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return database.DB.Create(user).Error
}

// func (r *userRepository) Login(email string, password string) (*models.User, error) {

// 	var user models.User

// 	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
// 		return nil, fmt.Errorf("user not found")
// 	}

// 	if !passwords.CompareHashPassword(password, user.Password) {
// 		return nil, fmt.Errorf("password not match")
// 	}
// 	return &user, nil

// }
