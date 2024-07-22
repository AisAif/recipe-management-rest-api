package models

import (
	"time"
)

type Recipe struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"size:255"`
	ImageURL  string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	IsPublic  bool
	Username  string `gorm:"size:255"`
	User      User   `gorm:"foreignKey:Username;references:Username"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
