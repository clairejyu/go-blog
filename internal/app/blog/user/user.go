package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andskur/argon2-hashing"
	"github.com/clairejyu/go-blog/internal/pkg/common"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName string `json:"nickName" binding:"required,min=1,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"-" binding:"required,min=6,max=255`
}

type Register struct {
	User
	Password string `json:"password" binding:"required,min=6,max=255`
}

func (u *Register) Create() *common.Message {
	_, err := GetByEmail(u.Email)
	if err == nil {
		return common.Fail(http.StatusBadRequest, "user exist. ")
	}

	// Generates a derived key with default params
	hash, err := argon2.GenerateFromPassword([]byte(u.Password), argon2.DefaultParams)
	if err != nil {
		return common.Fail(http.StatusInternalServerError, err.Error())
	}

	u.Password = string(hash)
	result := db.D.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("users").Create(&u).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return common.Fail(http.StatusInternalServerError, result.Error())
	}
	return common.Success("ok", u)
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

func GetByEmail(email string) (*User, error) {
	var user User
	result := db.D.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func ByEmail(email string) (interface{}, error) {
	var user User
	result := db.D.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) Update(id string) *common.Message {
	user := GetById(id)
	if user.Code != 200 {
		return common.Fail(user.Code, "the id of user had not found. "+user.Err)
	}

	if strings.TrimSpace(u.Password) != "" {
		// Generates a derived key with default params
		hash, err := argon2.GenerateFromPassword([]byte(u.Password), argon2.DefaultParams)
		if err != nil {
			return common.Fail(http.StatusInternalServerError, err.Error())
		}
		u.Password = string(hash)
	}

	result := db.D.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user.Obj).Updates(User{NickName: u.NickName, Email: u.Email}).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return common.Fail(http.StatusInternalServerError, result.Error())
	}
	return user
}

func (u *User) UpdatePassword(originPassword string, newPassword string) *common.Message {
	err := argon2.CompareHashAndPassword([]byte(u.Password), []byte(originPassword))
	if err != nil {
		return common.Fail(http.StatusBadRequest, "origin password not correct. "+err.Error())
	}

	// Generates a derived key with default params
	hash, err := argon2.GenerateFromPassword([]byte(newPassword), argon2.DefaultParams)
	if err != nil {
		return common.Fail(http.StatusInternalServerError, err.Error())
	}

	result := db.D.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&u).Update("password", hash).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return common.Fail(http.StatusInternalServerError, result.Error())
	}
	return common.Success("change succeed", nil)
}

func Delete(id string) *common.Message {
	user := GetById(id)
	if user.Code != 200 {
		return common.Fail(user.Code, "the id of user had not found. "+user.Err)
	}

	result := db.D.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&user.Obj, id).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return common.Fail(http.StatusInternalServerError, result.Error())
	}
	return user
}

func SignIn(email string, password string) *common.Message {
	user, err := GetByEmail(email)
	if err != nil {
		return common.Fail(http.StatusBadRequest, "user not found. "+err.Error())
	}

	// Uses the parameters from the existing derived key. Return an error if they don't match.
	err = argon2.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(password)
	if err != nil {
		return common.Fail(http.StatusBadRequest, "password not correct. "+err.Error())
	}
	return common.Success("login succeed", user)
}
