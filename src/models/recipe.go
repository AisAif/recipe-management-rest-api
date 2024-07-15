package models

import (
	"time"
)

type Recipe struct {
	ID        uint64
	Title     string `gorm:"size:255"`
	ImageURL  string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	Username  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
