package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model

	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}
