package model

type BilibiliUp struct {
	Mid      int64  `gorm:"primary_key;column:mid" json:"mid" form:"mid"`
	Name     string `gorm:"column:name" json:"name" form:"name"`
	LastTime int64  `gorm:"column:last_time" json:"last_time" form:"last_time"`
}
