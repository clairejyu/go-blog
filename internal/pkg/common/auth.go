package common

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionHelper(f func(email string) (interface{}, error)) func(c *gin.Context, callback func(user interface{})) {
	return func(c *gin.Context, callback func(user interface{})) {
		checkSession(c, f, callback)
	}
}

func checkSession(c *gin.Context, getUser func(email string) (interface{}, error), callback func(user interface{})) {
	session := sessions.Default(c)
	email := session.Get("email")
	if email != nil {
		user, err := getUser(email.(string))
		if err != nil {
			Fail(http.StatusInternalServerError, err.Error()).JSON(c)
		} else {
			callback(user)
		}
	} else {
		Fail(http.StatusForbidden, "need login").JSON(c)
	}
}
