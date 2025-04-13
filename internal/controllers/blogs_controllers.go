package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/services"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateBlogRequest struct {
	Title      string            `json:"title" binding:"required"`
	Content    string            `json:"content" binding:"required"`
	UserID     string            `json:"user_id" binding:"required"`
	Categories []models.Category `json:"categories" binding:"required"`
}

type GetBlogResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
	User      User   `json:"user"`
	Claps     int    `json:"claps"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateBlog godoc
// @Tags Blogs
// @Summary Create a new blog
// @Description Create a new blog
// @Accept  json
// @Produce  json
// @Param blog body CreateBlogRequest true "Blog object"
// @Success 201
// @Router /blog [post]
// @Security BearerAuth
func CreateBlog(c *gin.Context) {

	var req CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	user, err := services.GetUserService().GetById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var categories []models.Category
	for _, cat := range req.Categories {
		var existing models.Category
		if err := database.DB.First(&existing, "id = ?", cat.ID).Error; err == nil {
			categories = append(categories, existing)
		}
	}

	blog := &models.Blog{
		Title:      req.Title,
		Content:    req.Content,
		UserID:     userID,
		Categories: categories,
	}

	if err := services.GetBlogService().CreateBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, blog)

}

// GetBlog godoc
// @Tags Blogs
// @Summary Get a blog by ID
// @Description Get a blog by ID
// @ID get-blog-by-id
// @Produce  json
// @Param blogId path string true "Blog ID"
// @Success 200 {object} GetBlogResponse
// @Router /blog/{blogId} [get]
// @Security BearerAuth
func GetBlog(c *gin.Context) {
	id := c.Param("blogId")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		// Handle the error appropriately (maybe return a 400 Bad Request response)
	}

	blog, err := services.GetBlogService().GetBlogById(uuidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// GetAllBlogs godoc
// @Tags Blogs
// @Summary Retrieve all blogs
// @Description Retrieve a list of all blogs
// @Produce json
// @Success 200 {array} models.Blog
// @Router /blogs [get]
func GetAllBlogs(c *gin.Context) {
	blogs, err := services.GetBlogService().GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}
