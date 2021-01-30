package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	user.Create()
	c.JSON(http.StatusOK, user)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := GetById(id)
	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, List(c))
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	c.Bind(&user)
	c.JSON(http.StatusOK, user.Update(id))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, Delete(id))
}
