package blog

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	UserID  uint
	User    User   `json:"author"`
	Content string `json:"content"`
	State   int8   `json:"state"`
}
