package model

import (
	"Spider/database"
	"fmt"
	"github.com/jameskeane/bcrypt"
	"github.com/jinzhu/gorm"
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
func UserAdminCheckLogin(username string) User {
	u := User{}
	if err := database.DB.Where("username = ?", username).First(&u).Error; err != nil {
		fmt.Printf("UserAdminCheckLoginErr:%s", err)
	}
	return u
}

/**
 * 创建
 */
func CreateUser(user *User) (*User, error) {
	salt, _ := bcrypt.Salt(10)
	hash, _ := bcrypt.Hash(user.Password, salt)
	user.Password = string(hash)

	if err := database.DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return user, err
	}
	return user, nil
}

func FindUserByUsername(username string) (user *User) {
	if database.DB.Where("username = ?", username).First(&user).RecordNotFound() {
		return nil
	}
	return user
}
