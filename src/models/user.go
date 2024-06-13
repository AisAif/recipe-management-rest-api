package models

import "time"

type User struct {
	Username  string `gorm:"primaryKey"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
