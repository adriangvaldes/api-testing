package main

import (
	"github.com/adriangvaldes/api-testing/initializers"
	"github.com/adriangvaldes/api-testing/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
}
