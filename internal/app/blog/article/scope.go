package article

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func AuthorAndState(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		state := c.Query("state")
		userID := c.Query("userID")
		if strings.TrimSpace(state) != "" {
			db = db.Where("state = ?", state)
		}
		if strings.TrimSpace(userID) != "" {
			db = db.Where("user_id = ?", userID)
		}

		return db
	}
}
