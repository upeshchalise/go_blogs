package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/upeshchalise/go_blogs/internal/controllers"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/login", controllers.LoginUser)
	return r
}
