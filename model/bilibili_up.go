package model

import (
	"Spider/database"
	"time"
)

type BilibiliUp struct {
	ID            uint64          `gorm:"primary_key" mapstructure:"mid"`
	Name          string          `gorm:"column:name" json:"name"`
	LastTime      time.Time       `gorm:"column:last_time" json:"last_time"`
	Face          string          `gorm:"column:face" json:"face"`
	Sign          string          `gorm:"column:sign" json:"sign"`
	TopPhoto      string          `gorm:"column:top_photo" json:"top_photo" mapstructure:"top_photo"`
	BilibiliVideo []BilibiliVideo `gorm:"ForeignKey:Aid"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

//func New(id uint64, name string, face string, sign string, topPhoto string) (this *BilibiliUp) {
//	this = &BilibiliUp{}
//}

func FindBilibiliUpById(mid int64) *BilibiliUp {
	bilibiliUp := &BilibiliUp{}
	database.DB.First(&bilibiliUp, mid)
	return bilibiliUp
}
