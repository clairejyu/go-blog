package user

import "gorm.io/gorm"

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
