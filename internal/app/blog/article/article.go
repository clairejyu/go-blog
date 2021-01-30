package article

import (
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	UserID  uint
	User    user.User `json:"author"`
	Content string    `json:"content"`
	State   int8      `json:"state"`
}
