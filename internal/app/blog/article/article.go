package article

import (
	"net/http"

	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string    `json:"title" binding:"required,min=1,max=255"`
	UserID  uint      `json:"userID" binding:"required"`
	User    user.User `json:"author"`
	Content string    `json:"content"`
	State   int8      `json:"state" binding:"required"`
}

func (a *Article) Create() *common.Message {
	result := db.D.Create(&a)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return common.Success("ok", result)
}

func GetById(id string) *common.Message {
	var article Article
	result := db.D.Preload("User").First(&article, id)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return common.Success("ok", &article)
}

func List(c *gin.Context) *common.Message {
	var articles []*Article
	result := db.D.Scopes(db.Like(c, "title"), AuthorAndState(c), db.Paginate(c)).Preload("User").Order("created_at desc").Find(&articles)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return common.Success("ok", articles)
}

func (a *Article) Update(id string) *common.Message {
	article := GetById(id)
	if article.Code != 200 {
		return common.Fail(article.Code, "the id of article had not found. "+article.Err)
	}
	result := db.D.Model(article.Obj).Updates(&a)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return article
}

func Delete(id string) *common.Message {
	article := GetById(id)
	if article.Code != 200 {
		return common.Fail(article.Code, "the id of article had not found. "+article.Err)
	}
	result := db.D.Delete(&article.Obj, id)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return article
}
