package model

import (
	"time"
)

type BilibiliVideo struct {
	ID          uint64 `gorm:"primary_key" json:"aid"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	Mid         int    `gorm:"column:mid" json:"mid"`
	Title       string `gorm:"column:title" json:"title"`
	Created     uint64 `gorm:"column:created" json:"created"`
	Description string `gorm:"column:description" json:"description"`
	Pic         string `gorm:"column:pic" json:"pic"`
}
