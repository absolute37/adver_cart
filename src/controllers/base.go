package controllers

import (
	"adver_cart/src/repositories"

	"gorm.io/gorm"
)

type controllers struct {
}

func NewControllers(db *gorm.DB) *controllers {

	repo := repositories.BaseRepo{
		GromDB: db,
	}

	_ = repo

	return &controllers{}
}
