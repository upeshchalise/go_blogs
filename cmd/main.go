package main

import (
	"github.com/upeshchalise/go_blogs/config"
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/internal/routes"
	"github.com/upeshchalise/go_blogs/pkg/logger"

	_ "github.com/upeshchalise/go_blogs/docs"
)

// @title Go Blogs API
// @version 1.0
// @description This is a simple blog API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.ApiKey BearerAuth
// @in header
// @name Authorization
// @description Add token in "Bearer {token}“ format.
func main() {
	config.Load()
	logger.Init()

	database.Init()

	r := routes.InitRoutes()
	r.SetTrustedProxies([]string{})
	r.Run(":8080")
}
