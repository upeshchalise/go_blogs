package services

import (
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/repository"
)

type UserService interface {
	GetById(id uuid.UUID) (*models.User, error)
	Create(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func GetUserService() UserService {
	return &userService{
		repo: repository.NewUserRepository(),
	}
}

func (s *userService) GetById(id uuid.UUID) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) Create(user *models.User) error {
	return s.repo.Create(user)
}
