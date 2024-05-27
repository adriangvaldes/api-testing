package router

import (
	"github.com/adriangvaldes/api-testing/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(authController *controllers.AuthController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/auth")

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	return service
}
