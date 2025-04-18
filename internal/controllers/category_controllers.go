package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/services"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @ID create-category
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body CreateCategoryRequest true "Category object"
// @Param userId path string true "User ID"
// @Success 201 {object} models.Category
// @Router /category/user/{userId} [post]
// @Security BearerAuth
func CreateCategory(c *gin.Context) {
	var request CreateCategoryRequest
	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuidID, err := uuid.Parse(userId)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
	}

	user, err := services.GetUserService().GetById(uuidID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.UserType != models.UserRoles("admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	category := models.Category{
		Name: request.Name,
	}

	if err := services.GetCategoryService().CreateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)

}

// GetCategories godoc
// @Summary Retrieve all categories
// @Description Retrieve a list of all categories
// @Tags Categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	categories, err := services.GetCategoryService().GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}
