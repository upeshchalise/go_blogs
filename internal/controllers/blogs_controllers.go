package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/services"
)

type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

type GetBlogResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
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
// @Success 201 {object} GetBlogResponse
// @Router /blog [post]
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
		// Handle the error appropriately (maybe return a 400 Bad Request response)
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

	blog := &models.Blog{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := services.GetBlogService().CreateBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, blog)

}
