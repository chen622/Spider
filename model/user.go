package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nickname   string       `gorm:"column:nickname;size:64;not null" json:"nickname"`
	Password   string       `gorm:"column:password;size:255;not null" json:"password"`
	Mail       string       `gorm:"column:mail;size:128;not null" json:"mail"`
	BilibiliUp []BilibiliUp `gorm:"many2many:user_bilibili_up"`
}
