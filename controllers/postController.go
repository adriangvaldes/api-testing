package controllers

import (
	"github.com/adriangvaldes/api-testing/initializers"
	"github.com/adriangvaldes/api-testing/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body

	// Create a post
	post := models.Post{Title: "First post", Body: "Post body"}

	result := initializers.DB.Create(&post)
	// Return it
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
