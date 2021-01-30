package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article Article
	c.Bind(&article)
	article.Create()
	c.JSON(http.StatusOK, article)
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	article := GetById(id)
	c.JSON(http.StatusOK, article)
}

func ListArticles(c *gin.Context) {
	c.JSON(http.StatusOK, List(c))
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article Article
	c.Bind(&article)
	c.JSON(http.StatusOK, article.Update(id))
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, Delete(id))
}
