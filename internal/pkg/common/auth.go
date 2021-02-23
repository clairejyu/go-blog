package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NeedLogin(c *gin.Context, callback func(user interface{})) {
	currentUser, exist := c.Get("currentUser")
	if exist {
		callback(currentUser)
	}
	Fail(http.StatusForbidden, "need login").JSON(c)
}

func IsLogin(c *gin.Context) bool {
	_, exist := c.Get("currentUser")
	return exist
}
