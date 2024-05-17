package main

import (
	"github.com/adriangvaldes/api-testing/controllers"
	"github.com/adriangvaldes/api-testing/initializers"
	"github.com/gin-gonic/gin"
)

func init() { // INITIALIZERS
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.PostsCreate)
	r.Run()
}
