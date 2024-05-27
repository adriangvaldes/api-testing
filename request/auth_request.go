package request

import "gorm.io/gorm"

type RegisterRequest struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	gorm.Model
	Email    string
	Password string
}
