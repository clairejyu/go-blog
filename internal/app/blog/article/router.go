package article

import (
	"net/http"
	"strconv"

	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	common.NeedLogin(c, func(u interface{}) {
		currentUser := u.(*user.User)
		var article Article
		c.ShouldBind(&article)
		article.UserID = currentUser.ID
		article.Create().JSON(c)
	})
}

func GetArticleById(c *gin.Context) {
	GetById(c.Param("id")).JSON(c)
}

func ListArticles(c *gin.Context) {
	List(c).JSON(c)
}

func UpdateArticle(c *gin.Context) {
	common.NeedLogin(c, func(u interface{}) {
		currentUser := u.(*user.User)
		id := c.Param("id")
		if id == strconv.FormatUint(uint64(currentUser.ID), 10) {
			var article Article
			c.ShouldBind(&article)
			article.Update(id).JSON(c)
		} else {
			common.Fail(http.StatusForbidden, "need login").JSON(c)
		}
	})
}

func DeleteArticle(c *gin.Context) {
	common.NeedLogin(c, func(u interface{}) {
		currentUser := u.(*user.User)
		id := c.Param("id")
		if id == strconv.FormatUint(uint64(currentUser.ID), 10) {
			Delete(id).JSON(c)
		} else {
			common.Fail(http.StatusForbidden, "need login").JSON(c)
		}
	})
}
