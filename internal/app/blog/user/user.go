package user

import (
	"net/http"

	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName string `json:"nickName" binding:"required,min=1,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=255`
}

func (u *User) Create() *common.Message {
	result := db.D.Create(&u)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return common.Success("ok", result)
}

func GetById(id string) *common.Message {
	var user User
	result := db.D.First(&user, id)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return common.Success("ok", &user)
}

func List(c *gin.Context) *common.Message {
	var users []*User
	result := db.D.Scopes(db.Like(c, "nickName"), db.Like(c, "email"), db.Paginate(c)).Order("nick_name asc").Find(&users)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return common.Success("ok", users)
}

func (u *User) ListUsersByNickNameAndEmail() []*User {
	var users []*User
	db.D.Where(&u).Find(&users)
	return users
}

func (u *User) Update(id string) *common.Message {
	user := GetById(id)
	if user.Code != 200 {
		return common.Fail(user.Code, "the id of user had not found. "+user.Err)
	}

	result := db.D.Model(user).Updates(&u)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return user
}

func Delete(id string) *common.Message {
	user := GetById(id)
	if user.Code != 200 {
		return common.Fail(user.Code, "the id of user had not found. "+user.Err)
	}

	result := db.D.Delete(&user, id)
	if result.Error != nil {
		return common.Fail(http.StatusInternalServerError, result.Error.Error())
	}
	return user
}
