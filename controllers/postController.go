package controllers

import "github.com/gin-gonic/gin"

func PostsCreate(c *gin.Context) {
	// Get data off req body

	// Create a post

	// Return it

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
