package main

import (
	"github.com/clairejyu/go-blog/cmd/blog/router"
	"github.com/clairejyu/go-blog/internal/app/blog/article"
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func init() {
	db.Init()
	user.InitDB(db.D)
	article.InitDB(db.D)
}

func main() {
	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("hogwarts"))
	r.Use(sessions.Sessions("user", store))

	v1 := r.Group("/api")
	router.Auth(v1)
	router.User(v1.Group("/user"))
	router.Article(v1.Group("/article"))
	r.Run()
}
