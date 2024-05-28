package main

import (
	"github.com/adriangvaldes/api-testing/controllers"
	"github.com/adriangvaldes/api-testing/initializers"
	"github.com/adriangvaldes/api-testing/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() { // INITIALIZERS TO LOAD
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// POSTS ROUTES
	postRoutes := r.Group("/posts")
	postRoutes.Use(middlewares.JwtAuthMiddleware())
	postRoutes.POST("/", controllers.PostsCreate)
	postRoutes.GET("/", controllers.PostIndex)
	postRoutes.GET("/:id", controllers.PostsShow)
	postRoutes.PUT("/:id", controllers.PostsUpdate)
	postRoutes.DELETE("/:id", controllers.PostsDelete)

	validate := validator.New()
	authController := controllers.NewAuthControllerImpl(initializers.DB, validate)

	// AUTH ROUTES
	authRoutes := r.Group("/auth")
	authRoutes.POST("/register", authController.Register)
	authRoutes.POST("/login", authController.Login)

	r.Run()

}
