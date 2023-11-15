package db

import (
	"fmt"

	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	"github.com/Nishad4140/ecommerce_project/routines"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabse(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", cfg.DBHost, cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSslmode)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&domain.Users{},
		&domain.Admins{},
		&domain.SupAdmins{},
		&domain.UserBlockInfo{},
		&domain.UserReportInfo{},
		&domain.Address{},
		&domain.Category{},
		&domain.Brands{},
		&domain.Model{},
		&domain.Carts{},
		&domain.CartItem{},
		&domain.Orders{},
		&domain.OrderItem{},
		&domain.OrderStatus{},
		&domain.PaymentType{},
		&domain.PaymentDetails{},
		&domain.PaymentStatus{},
		&domain.Images{},
		&domain.UserWallet{},
	)

	routines := routines.NewConcurrency(db)

	routines.GetConcurrency()

	return db, dbErr
}
