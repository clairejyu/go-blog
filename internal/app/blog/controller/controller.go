package controller

import (
	"fmt"
	"net/http"

	"github.com/clairejyu/go-blog/internal/app/blog"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user blog.User
	c.Bind(&user)
	user.Create()
	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := blog.GetUserById(id)
	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, blog.ListUsers(c))
}

// func ListUsersByNickName(c *gin.Context) {
// 	c.JSON(http.StatusOK, blog.ListUsersByNickName(c.Param("nickName")))
// }
