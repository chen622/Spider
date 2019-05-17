package model

import (
	"time"
)

type BilibiliVideo struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Aid         int        `gorm:"column:aid" json:"aid"`
	Title       string     `gorm:"column:title" json:"title"`
	Created     time.Time  `gorm:"column:created" json:"created"`
	Description string     `gorm:"column:description" json:"description"`
	Pic         string     `gorm:"column:pic" json:"pic"`
}
