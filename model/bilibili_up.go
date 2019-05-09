package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BilibiliUp struct {
	gorm.Model
	Name          string          `gorm:"column:name" json:"name"`
	LastTime      time.Time       `gorm:"column:last_time" json:"last_time"`
	Face          string          `gorm:"column:face" json:"face"`
	Sign          string          `gorm:"column:sign" json:"sign"`
	TopPhoto      string          `gorm:"column:top_photo" json:"top_photo"`
	BilibiliVideo []BilibiliVideo `gorm:"ForeignKey:Aid"`
}
