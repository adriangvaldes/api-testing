package main

import (
	"github.com/adriangvaldes/api-testing/controllers"
	"github.com/adriangvaldes/api-testing/initializers"
	"github.com/gin-gonic/gin"
)

func init() { // INITIALIZERS TO LOAD
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
