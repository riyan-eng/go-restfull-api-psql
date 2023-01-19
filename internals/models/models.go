package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Title    string `json:"Title" validate:"required"`
	SubTitle string `json:"SubTitle" validate:"required"`
	Text     string `json:"Text" validate:"required"`
}
