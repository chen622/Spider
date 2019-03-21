package model

type User struct {
	ID       int64  `gorm:"primary_key;column:id" json:"id" form:"id"`
	Nickname string `gorm:"column:nickname" json:"nickname" form:"nickname"`
	Password string `gorm:"column:password" json:"password" form:"password"`
	Mail     string `gorm:"column:mail" json:"mail" form:"mail"`
}
