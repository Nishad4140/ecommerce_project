package domain

import "time"

type Users struct {
	ID        uint `gorm:"primaryKey;unique;not null"`
	Name      string
	Email     string `gorm:"unique;not null"`
	Mobile    string `gorm:"unique; not null"`
	Password  string `gorm:"not null"`
	IsBlocked bool   `gorm:"default:false"`
	CreatedAt time.Time
}
