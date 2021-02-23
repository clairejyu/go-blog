package main

import (
	"net/http"

	"github.com/clairejyu/go-blog/cmd/blog/router"
	"github.com/clairejyu/go-blog/internal/app/blog/article"
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/clairejyu/go-blog/internal/pkg/ginext"
	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:8080", "https://waita.site"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("hogwarts"))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24 * 30,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("user", store))
	r.Use(ginext.SetCurrentUser())
	v1 := r.Group("/api")
	router.Auth(v1)
	router.User(v1.Group("/user"))
	router.Article(v1.Group("/article"))
	r.Run()
}
