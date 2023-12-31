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
	BlockUntil        time.Time
	ReasonForBlocking string
}

type Address struct {
	ID           uint `gorm:"primaryKey;unique;not null"`
	UsersID      uint
	Users        Users  `gorm:"foreignKey:UsersID"`
	House_number string `json:"house_number" binding:"required"`
	Street       string `json:"street" binding:"required"`
	City         string `json:"city " binding:"required"`
	District     string `json:"district " binding:"required"`
	Landmark     string `json:"landmark" binding:"required"`
	Pincode      int    `json:"pincode " binding:"required"`
	IsDefault    bool   `gorm:"default:false"`
}

type UserWallet struct {
	Id      uint `gorm:"primaryKey"`
	UsersId uint
	Users   Users `gorm:"foreignKey:UsersId"`
	Amount  int   `gorm:"default:0;check:Amount>=0" sql:"CHECK(Amount >= 0)"`
	IsLock  bool  `gorm:"default:true"`
}
