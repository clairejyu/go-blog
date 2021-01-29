package main

import (
	"github.com/clairejyu/go-blog/internal/app/blog"
	"github.com/clairejyu/go-blog/internal/app/blog/controller"
	"github.com/clairejyu/go-blog/internal/pkg"
	"github.com/clairejyu/go-blog/internal/pkg/setting"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	setting.Setup()
	db := pkg.InitDB()
	Migrator(db)
}

func main() {
	r := gin.Default()
	r.POST("/user", controller.CreateUser)
	r.GET("/user/:id", controller.GetUserById)
	r.GET("/users", controller.ListUsers)
	//r.GET("/user/:id", controller.ListUsersByNickName)
	r.Run()
}

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&blog.User{}, &blog.Article{})
}
