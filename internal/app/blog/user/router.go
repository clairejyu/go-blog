package user

import (
	"net/http"
	"strconv"

	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

var session = common.SessionHelper(ByEmail)

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
	session(c, func(u interface{}) {
		currentUser := u.(User)
		id := c.Param("id")
		if id == strconv.FormatUint(uint64(currentUser.ID), 10) {
			var user User
			c.ShouldBind(&user)
			user.Update(id).JSON(c)
		} else {
			common.Fail(http.StatusForbidden, "need login").JSON(c)
		}
	})
}

func DeleteUser(c *gin.Context) {
	Delete(c.Param("id")).JSON(c)
}

func Login(c *gin.Context) {
	email := c.Query("email")
	result := SignIn(email, c.Query("password"))
	if result.Code == 200 {
		session := sessions.Default(c)
		session.Set("email", email)
		session.Save()
	}
	result.JSON(c)
}

func ChangePassword(c *gin.Context) {
	session(c, func(u interface{}) {
		user := u.(User)
		user.UpdatePassword(c.Query("originPassword"), c.Query("newPassword")).JSON(c)
	})
}

func Exit(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("email")
	session.Save()
	c.JSON(http.StatusOK, "Exit succeed")
}
