package user

import (
	"net/http"
	"strconv"

	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var register Register
	err := c.ShouldBind(&register)
	if err != nil {
		common.Fail(http.StatusBadRequest, err.Error()).JSON(c)
	} else {
		register.Create().JSON(c)
	}
}

func GetUserById(c *gin.Context) {
	GetById(c.Param("id")).JSON(c)
}

func ListUsers(c *gin.Context) {
	List(c).JSON(c)
}

func UpdateUser(c *gin.Context) {
	common.NeedLogin(c, func(u interface{}) {
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

type LoginParam struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=255`
}

func Login(c *gin.Context) {
	var param LoginParam
	err := c.ShouldBind(&param)
	if err != nil {
		common.Fail(http.StatusBadRequest, err.Error()).JSON(c)
	} else {
		email := param.Email
		result := SignIn(param.Email, param.Password)
		if result.Code == 200 {
			session := sessions.Default(c)
			session.Set("email", email)
			session.Save()
		}
		result.JSON(c)
	}
}

func ChangePassword(c *gin.Context) {
	common.NeedLogin(c, func(u interface{}) {
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
