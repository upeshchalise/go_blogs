package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/services"
	"github.com/upeshchalise/go_blogs/pkg/utils/jwt"
	passwords "github.com/upeshchalise/go_blogs/pkg/utils/password"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// GetUser godoc
// @Tags Users
// @Summary Get a user by ID
// @Description Get a user by ID
// @ID get-user-by-id
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
// @Security BearerAuth
func GetUser(c *gin.Context) {
	id := c.Param("id")

	uuidID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Invalid UUID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		// Handle the error appropriately (maybe return a 400 Bad Request response)
	}

	user, err := services.GetUserService().GetById(uuidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)

}

// CreateUser godoc
// @Tags Users
// @Summary Create a new user
// @Description Create a new user
// @ID create-user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User object"
// @Success 201 {object} models.User
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := passwords.HashPassword(string(user.Password))
	if err != nil {
		log.Println("error while hashing the password")
	}

	user.Password = hashedPassword

	if err := services.GetUserService().Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{})
}

// LoginUser godoc
// @Tags Users
// @Summary Login a user
// @Description Login a user
// @ID login-user
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User object"
// @Success 200 {object} models.User
// @Router /login [post]
func LoginUser(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	fmt.Println(req.Email, req.Password)
	user, err := services.GetUserService().GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateJwtToken(user.ID.String(), "thisissecrettoken")

	if err != nil {
		log.Println("error while generating jwt token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}
