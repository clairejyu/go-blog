package router

import (
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	router.PUT("/", user.CreateUser)
	router.GET("/:id", user.GetUserById)
	router.GET("/", user.ListUsers)
	router.POST("/:id", user.UpdateUser)
	router.DELETE("/:id", user.DeleteUser)
}

func Article(router *gin.RouterGroup) {

}
