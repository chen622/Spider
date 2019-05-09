package model

type BilibiliUp struct {
	Mid      int64  `jpath:"mid" gorm:"primary_key;column:mid" json:"mid" form:"mid"`
	Name     string `jpath:"name" gorm:"column:name" json:"name" form:"name"`
	LastTime int64  `jpath:"" gorm:"column:last_time" json:"last_time" form:"last_time"`
	Face     string `jpath:"" gorm:"column:face" json:"face" form:"face"`
	Sign     string `jpath:"" gorm:"column:sign" json:"sign" form:"sign"`
	TopPhoto string `gorm:"column:top_photo" json:"top_photo" form:"top_photo"`
}
