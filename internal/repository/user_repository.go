package repository

import (
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/models"
)

type UserRepository interface {
	GetById(id uuid.UUID) (*models.User, error)
	Create(user *models.User) error
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

func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
