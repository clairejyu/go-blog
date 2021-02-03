package router

import (
	"github.com/clairejyu/go-blog/internal/app/blog/article"
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	router.POST("/login", user.Login)
	router.POST("/change-password", user.ChangePassword)
	router.DELETE("/exit", user.Exit)
}

func User(router *gin.RouterGroup) {
	router.PUT("/", user.CreateUser)
	router.GET("/:id", user.GetUserById)
	router.GET("/", user.ListUsers)
	router.POST("/:id", user.UpdateUser)
	router.DELETE("/:id", user.DeleteUser)
}

func Article(router *gin.RouterGroup) {
	router.PUT("/", article.CreateArticle)
	router.GET("/:id", article.GetArticleById)
	router.GET("/", article.ListArticles)
	router.POST("/:id", article.UpdateArticle)
	router.DELETE("/:id", article.DeleteArticle)
}
