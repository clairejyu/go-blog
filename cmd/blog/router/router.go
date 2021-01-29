package router

import (
	_ "net/http"

	_ "github.com/clairejyu/go-blog/internal/app/blog"
	"github.com/gin-gonic/gin"
)

func UserGet(c *gin.Context) {
	// c.JSON(http.StatusOK, blog.GetUser())
}
