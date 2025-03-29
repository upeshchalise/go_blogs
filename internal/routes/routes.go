package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/upeshchalise/go_blogs/internal/controllers"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	apiV1.POST("/user", controllers.CreateUser)
	apiV1.GET("/user/:id", controllers.GetUser)
	apiV1.POST("/login", controllers.LoginUser)
	apiV1.POST("/blog", controllers.CreateBlog)
	apiV1.GET("/blog/:blogId", controllers.GetBlog)
	return r
}
