package services

import (
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/repository"
)

type BlogService interface {
	CreateBlog(blog *models.Blog) error
	GetBlogById(id uuid.UUID) (*models.Blog, error)
	GetAllBlogs() ([]models.Blog, error)
	GetBlogsByUserId(userId uuid.UUID) ([]models.Blog, error)
}

type blogService struct {
	repo repository.BlogsRepository
}

func GetBlogService() BlogService {
	return &blogService{
		repo: repository.NewBlogsRepository(),
	}
}

func (s *blogService) CreateBlog(blog *models.Blog) error {
	return s.repo.CreateBlog(blog)
}

func (s *blogService) GetBlogById(id uuid.UUID) (*models.Blog, error) {
	return s.repo.GetBlogById(id)
}

func (s *blogService) GetAllBlogs() ([]models.Blog, error) {
	return s.repo.GetAllBlogs()
}

func (s *blogService) GetBlogsByUserId(userId uuid.UUID) ([]models.Blog, error) {
	return s.repo.GetBlogsByUserId(userId)
}
