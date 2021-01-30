package main

import (
	"github.com/clairejyu/go-blog/cmd/blog/router"
	"github.com/clairejyu/go-blog/internal/app/blog/article"
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	db.Setup()
	db := db.InitDB()
	Migrator(db)
}

func main() {
	r := gin.Default()
	v1 := r.Group("/api")
	router.User(v1.Group("/user"))
	router.Article(v1.Group("/article"))
	r.Run()
}

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &article.Article{})
}
