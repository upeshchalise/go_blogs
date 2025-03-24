package repository

import (
	"fmt"

	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/models"
)

type BlogsRepository interface {
	CreateBlog(blog *models.Blog) error
	// GetBlogById(id uuid.UUID) (*models.Blog, error)
	// GetAllBlogs() ([]models.Blog, error)
	// GetBlogsByUserId(userId uuid.UUID) ([]models.Blog, error)
}

type blogsRepository struct{}

func NewBlogsRepository() BlogsRepository {
	return &blogsRepository{}
}

func (r *blogsRepository) CreateBlog(blog *models.Blog) error {

	if database.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return database.DB.Create(blog).Error
}
