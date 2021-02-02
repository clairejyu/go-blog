package user

import (
	"fmt"
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

func Login(c *gin.Context) {
	fmt.Println("c ", c)
	SignIn(c.Query("email"), c.Query("password")).JSON(c)
}

func ChangePassword(c *gin.Context) {
	email := c.Query("email")
	user, err := GetByEmail(email)
	if err != nil {
		common.Fail(http.StatusInternalServerError, err.Error()).JSON(c)
	} else {
		user.UpdatePassword(c.Query("originPassword"), c.Query("newPassword")).JSON(c)
	}
}
