package article

import (
	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string    `json:"title"`
	UserID  uint      `json:"userID"`
	User    user.User `json:"author"`
	Content string    `json:"content"`
	State   int8      `json:"state"`
}

func (a *Article) Create() *gorm.DB {
	return db.D.Create(&a)
}

func GetById(id string) *Article {
	var article Article
	db.D.Preload("User").First(&article, id)
	return &article
}

func List(c *gin.Context) []*Article {
	var articles []*Article
	d := db.D
	d.Scopes(db.Like(c, "title"), AuthorAndState(c), db.Paginate(c)).Preload("User").Order("created_at desc").Find(&articles)
	return articles
}

func (a *Article) Update(id string) *Article {
	article := GetById(id)
	if article != nil {
		db.D.Model(article).Updates(&a)
	}
	return article
}

func Delete(id string) *Article {
	article := GetById(id)
	db.D.Delete(&article, id)
	return article
}
