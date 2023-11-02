package repository

import (
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	"gorm.io/gorm"
)

type PaymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &PaymentDatabase{DB}
}
