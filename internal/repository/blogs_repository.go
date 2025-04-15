package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/models"
	"gorm.io/gorm"
)

type BlogsRepository interface {
	CreateBlog(blog *models.Blog) error
	GetBlogById(id uuid.UUID) (*models.Blog, error)
	GetAllBlogs() ([]models.Blog, error)
	GetBlogsByUserId(userId uuid.UUID) ([]models.Blog, error)
	GetBlogsByCategory(categoryId uuid.UUID) ([]models.Blog, error)
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

func (r *blogsRepository) GetBlogById(id uuid.UUID) (*models.Blog, error) {
	var blog models.Blog
	if err := database.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "first_name", "last_name") }).
		Preload("Categories", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		First(&blog, id).
		Error; err != nil {

		return nil, err
	}

	return &blog, nil
}

func (r *blogsRepository) GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := database.DB.
		Preload("Categories", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *blogsRepository) GetBlogsByUserId(userId uuid.UUID) ([]models.Blog, error) {
	var blogs []models.Blog
	if err := database.DB.Where("user_id = ?", userId).Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *blogsRepository) GetBlogsByCategory(categoryId uuid.UUID) ([]models.Blog, error) {
	var blogs []models.Blog

	err := database.DB.Joins("JOIN blog_categories ON blog_categories.blog_id = blogs.id").
		Where("blog_categories.category_id = ?", categoryId).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "first_name", "last_name") }).Preload("Categories").Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	return blogs, nil
}
