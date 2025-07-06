package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	PostID uint
	Post   Post `gorm:"constraint:OnDelete:CASCADE"`
	UserID string
	IsUp   bool
}