package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gotut/docs"
	"gotut/handlers"
	"gotut/middleware"
)

// @title Todo API
// @version 1.0
// @description This is a simple Todo API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host gotut-api.germanywestcentral.azurecontainer.io:8080
// @BasePath /
// @schemes http
func main() {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware())

	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id", handlers.GetTodoByID)
	router.POST("/todos", handlers.CreateTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen on all interfaces (0.0.0.0) for Azure Container Instances
	router.Run("0.0.0.0:8080")
}
