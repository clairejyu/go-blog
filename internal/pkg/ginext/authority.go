package ginext

import (
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")
		if email != nil {
			user, err := user.GetByEmail(email.(string))
			if err == nil {
				c.Set("currentUser", user)
			}
		}
		c.Next()
	}
}
