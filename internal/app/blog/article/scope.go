package article

import (
	"strconv"
	"strings"

	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthorAndState(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		state := strings.TrimSpace(c.Query("state"))
		userID := strings.TrimSpace(c.Query("userID"))
		currentUser, exist := c.Get("currentUser")

		if state != "" && state == strconv.Itoa(2) && exist {
			db = db.Where("user_id = ?", currentUser.(*user.User).ID).Where("state = ?", 2)
		} else {
			db = db.Where("state = ?", 1)
		}

		if userID != "" {
			db = db.Where("user_id = ?", userID)
		}
		return db
	}
}
