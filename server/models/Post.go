package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content   string `json:"content"`
	Upvotes   int32  `json:"upvotes" gorm:"default:0"`
	Downvotes int32  `json:"downvotes" gorm:"default:0"`
	UserId string `json:"user_id"`
	Votes     []Vote `gorm:"foreignKey:PostID"`
}