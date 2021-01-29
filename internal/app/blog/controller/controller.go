package controller

import (
	"net/http"

	"github.com/clairejyu/go-blog/internal/app/blog"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user blog.User
	c.Bind(&user)
	user.Create()
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

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user blog.User
	c.Bind(&user)
	c.JSON(http.StatusOK, user.UpdateUser(id))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, blog.DeleteUser(id))
}
