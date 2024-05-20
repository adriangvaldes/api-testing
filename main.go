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
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
