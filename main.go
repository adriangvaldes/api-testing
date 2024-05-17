package main

import (
	"github.com/adriangvaldes/api-testing/initializers"
	"github.com/gin-gonic/gin"
)

func init() { // INITIALIZERS
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
