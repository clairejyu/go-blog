package user

import (
	"github.com/clairejyu/go-blog/internal/pkg/db"
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
	return db.D.Create(&u)
}

func GetById(id string) *User {
	var user User
	db.D.First(&user, id)
	return &user
}

func List(c *gin.Context) []*User {
	var users []*User
	d := db.D
	d.Scopes(db.Like(c, "nickName"), db.Like(c, "email"), db.Paginate(c)).Order("nick_name asc").Find(&users)
	return users
}

func (u *User) ListUsersByNickNameAndEmail() []*User {
	var users []*User
	db.D.Where(&u).Find(&users)
	return users
}

func (u *User) Update(id string) *User {
	user := GetById(id)
	if user != nil {
		db.D.Model(user).Updates(&u)
	}
	return user
}

func Delete(id string) *User {
	user := GetById(id)
	db.D.Delete(&user, id)
	return user
}
