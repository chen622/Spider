package model

import (
	"Spider/database"
	"fmt"
	"time"
)

type BilibiliUp struct {
	ID            uint64          `gorm:"primary_key" mapstructure:"mid"`
	Name          string          `gorm:"column:name" json:"name"`
	LastTime      time.Time       `gorm:"column:last_time" json:"lastTime"`
	Face          string          `gorm:"column:face" json:"face"`
	Sign          string          `gorm:"column:sign" json:"sign"`
	TopPhoto      string          `gorm:"column:top_photo" json:"topPhoto" mapstructure:"top_photo"`
	BilibiliVideo []BilibiliVideo `gorm:"ForeignKey:Aid"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

//func New(id uint64, name string, face string, sign string, topPhoto string) (this *BilibiliUp) {
//	this = &BilibiliUp{}
//}

func FindBilibiliUpById(mid int64) (*BilibiliUp, error) {
	bilibiliUp := &BilibiliUp{}
	database.DB.First(&bilibiliUp, mid)
	if bilibiliUp.ID == 0 {
		return nil, fmt.Errorf("No this up")
	} else {
		return bilibiliUp, nil
	}
}
