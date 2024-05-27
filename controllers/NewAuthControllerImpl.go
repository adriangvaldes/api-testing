package controllers

import (
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewAuthControllerImpl(Db *gorm.DB, validate *validator.Validate) *AuthController {
	return &AuthController{Db: Db, Validate: validate}
}
