package user

import (
	"net/http"

	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		common.Fail(http.StatusBadRequest, err.Error()).JSON(c)
	} else {
		user.Create().JSON(c)
	}
}

func GetUserById(c *gin.Context) {
	GetById(c.Param("id")).JSON(c)
}

func ListUsers(c *gin.Context) {
	List(c).JSON(c)
}

func UpdateUser(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	user.Update(c.Param("id")).JSON(c)
}

func DeleteUser(c *gin.Context) {
	Delete(c.Param("id")).JSON(c)
}
