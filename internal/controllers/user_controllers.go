package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upeshchalise/go_blogs/internal/models"
	"github.com/upeshchalise/go_blogs/internal/services"
)

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
	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(user.Password)

	if err := services.GetUserService().Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
