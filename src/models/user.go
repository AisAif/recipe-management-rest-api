package models

import "time"

type User struct {
	Username  string `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Password  string `gorm:"size:60"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
