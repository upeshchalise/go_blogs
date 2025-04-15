package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/upeshchalise/go_blogs/internal/controllers"
	"github.com/upeshchalise/go_blogs/internal/middleware"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CorsMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	apiV1.POST("/user", controllers.CreateUser)
	apiV1.GET("/user/:id", middleware.VerifyJWT(controllers.GetUser))
	apiV1.POST("/login", controllers.LoginUser)
	apiV1.POST("/blog", middleware.VerifyJWT(controllers.CreateBlog))
	apiV1.GET("/blogs", controllers.GetAllBlogs)
	apiV1.GET("/blog/:blogId", controllers.GetBlog)
	apiV1.GET("/blogs/category/:categoryId", controllers.GetBlogsByCategory)
	apiV1.POST("/category/user/:userId", middleware.VerifyJWT(controllers.CreateCategory))
	apiV1.GET("/categories", controllers.GetCategories)
	return r
}
