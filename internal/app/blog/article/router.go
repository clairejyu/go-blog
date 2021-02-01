package article

import (
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article Article
	c.ShouldBind(&article)
	article.Create().JSON(c)
}

func GetArticleById(c *gin.Context) {
	GetById(c.Param("id")).JSON(c)
}

func ListArticles(c *gin.Context) {
	List(c).JSON(c)
}

func UpdateArticle(c *gin.Context) {
	var article Article
	c.ShouldBind(&article)
	article.Update(c.Param("id")).JSON(c)
}

func DeleteArticle(c *gin.Context) {
	Delete(c.Param("id")).JSON(c)
}
