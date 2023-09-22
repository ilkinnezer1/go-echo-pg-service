package models

import "time"

type User struct {
	ID            uint   `gorm:"primary_key"`
	Username      string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	PrevLoginTime time.Time
	LastLoginTime time.Time
}
