package article

import "gorm.io/gorm"

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&Article{})
}