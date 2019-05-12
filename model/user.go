package model

import (
	"Spider/database"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username   string       `gorm:"column:username;size:64;not null;unique" json:"username" validate:"required,gte=4,lte=64"`
	Password   string       `gorm:"column:password;size:255;not null" json:"password" validate:"required,gte=7,lte=256"`
	Mail       string       `gorm:"column:mail;size:128;not null" json:"email" validate:"email"`
	BilibiliUp []BilibiliUp `gorm:"many2many:user_bilibili_up"`
}

/**
 * 校验用户登录
 * @method UserAdminCheckLogin
 * @param  {[type]}       username string [description]
 */
func UserAdminCheckLogin(username string) (*User, error) {
	u := &User{}
	if err := database.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

/**
 * 创建
 */
func CreateUser(user *User) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)
	if err := database.DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return user, err
	}
	return user, nil
}

func FindUserByUsername(username string) *User {
	user := &User{}
	database.DB.Where("username = ?", username).First(&user)
	return user
}

func FindUserById(id uint) *User {
	user := &User{}
	database.DB.First(&user, id)
	return user
}
