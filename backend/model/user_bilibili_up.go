package model

type UserBilibiliUp struct {
	UserId       int64 `gorm:"primary_key;column:user_id" json:"user_id" form:"user_id"`
	BilibiliUpId int64 `gorm:"column:bilibili_up_id" json:"bilibili_up_id" form:"bilibili_up_id"`
}
