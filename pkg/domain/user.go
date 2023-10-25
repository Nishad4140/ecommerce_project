package domain

import "time"

type Users struct {
	ID          uint   `gorm:"primaryKey;unique;not null"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Mobile      string `json:"mobile" binding:"required,eq=10" gorm:"unique; not null"`
	Password    string `json:"password" gorm:"not null"`
	ReportCount int    `gorm:"default:0"`
	IsBlocked   bool   `gorm:"default:false"`
	CreatedAt   time.Time
}

type UserReportInfo struct {
	ID                 uint `gorm:"primaryKey"`
	UsersID            uint
	Users              Users `gorm:"foreignKey:UsersID"`
	ReportedAt         time.Time
	ReportedBy         uint
	ReasonForReporting string
}

type UserBlockInfo struct {
	ID                uint `gorm:"primaryKey"`
	UsersID           uint
	Users             Users `gorm:"foreignKey:UsersID"`
	BlockedAt         time.Time
	BlockedBy         uint
	ReasonForBlocking string
}
