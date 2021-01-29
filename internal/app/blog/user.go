package blog

import (
	"github.com/clairejyu/go-blog/internal/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Create() *gorm.DB {
	return pkg.DB.Create(&u)
}

func GetUserById(id string) *User {
	var user User
	pkg.DB.First(&user, id)
	return &user
}

func ListUsers(c *gin.Context) []*User {
	var users []*User
	db := pkg.DB
	db.Scopes(pkg.Like(c, "nickName"), pkg.Like(c, "email"), pkg.Paginate(c)).Order("nick_name asc").Find(&users)
	return users
}

func (u *User) ListUsersByNickNameAndEmail() []*User {
	var users []*User
	pkg.DB.Where(&u).Find(&users)
	return users
}
