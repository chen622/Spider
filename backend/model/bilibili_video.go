package model

import (
	"time"
)

type BilibiliVideo struct {
	Mid         int    `jpath:"mid" gorm:"primary_key;column:mid" json:"mid" form:"mid"`
	Aid         int    `jpath:"aid" gorm:"column:aid" json:"aid" form:"aid"`
	Title       string `jpath:"title" gorm:"column:title" json:"title" form:"title"`
	Created     int64  `jpath:"created" gorm:"column:created" json:"created" form:"created"`
	Description string `jpath:"description" gorm:"column:description" json:"description" form:"description"`
	Pic         string `jpath:"pic" gorm:"column:pic" json:"pic" form:"pic"`
}

func (v BilibiliVideo) GetTime() time.Time {
	unix := time.Unix(v.Created, 0)
	return unix
}
