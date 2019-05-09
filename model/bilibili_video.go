package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BilibiliVideo struct {
	gorm.Model
	Aid         int       `gorm:"column:aid" json:"aid"`
	Title       string    `gorm:"column:title" json:"title"`
	Created     time.Time `gorm:"column:created" json:"created"`
	Description string    `gorm:"column:description" json:"description"`
	Pic         string    `gorm:"column:pic" json:"pic"`
}
