package repository

import (
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategories() ([]models.Category, error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (c *categoryRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categories, nil
}

func (c *categoryRepository) CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}
