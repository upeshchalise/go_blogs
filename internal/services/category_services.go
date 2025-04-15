package services

import (
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetCategories() ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func GetCategoryService() CategoryService {
	return &categoryService{
		repo: repository.NewCategoryRepository(),
	}
}

func (s *categoryService) GetCategories() ([]models.Category, error) {
	return s.repo.GetCategories()
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.CreateCategory(category)
}
