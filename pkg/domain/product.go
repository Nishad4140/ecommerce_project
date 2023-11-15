package domain

import "time"

type Category struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	CategoryName string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}

type Brands struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	BrandName   string `gorm:"unique;not null"`
	Description string
	Category_id uint
	Category    Category `gorm:"foreignKey:Category_id"`
	Created_at  time.Time
	Updated_at  time.Time
}

type Model struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	ModelName    string `gorm:"not null"`
	Brand_id     uint
	Brands       Brands `gorm:"foreignKey:Brand_id"`
	Sku          string `gorm:"not null"`
	Qty_in_stock int
	Color        string
	Ram          int
	Battery      int
	Screen_size  float64
	Storage      int
	Camera       int
	Price        int //`gorm:"not null;check:price>=0" sql:"CHECK(price >= 0)"`
	Image      string
	Created_at time.Time
	Updated_at time.Time
}

type Images struct {
	Id       uint `gorm:"primaryKey;unique;not null"`
	ModelId  uint
	Model    Model `gorm:"foreignKey:ModelId"`
	FileName string
}
