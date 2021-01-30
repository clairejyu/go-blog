package db

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

func Like(c *gin.Context, key string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		value := c.Query(key)
		if strings.TrimSpace(value) != "" {
			db = db.Where(strcase.ToSnake(key)+" LIKE ?", "%"+value+"%")
		}
		return db
	}
}
