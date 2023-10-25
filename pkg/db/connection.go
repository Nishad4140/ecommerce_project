package db

import (
	"fmt"

	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabse(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s", cfg.DBHost, cfg.DBName, cfg.DBUser, cfg.DBPassword, cfg.DBSslmode)
	// psqlInfo := "host=" + cfg.DBHost + " dbname=" + cfg.DBName + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " sslmode=" + cfg.DBSslmode
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&domain.Users{},
		&domain.Admins{},
		&domain.UserBlockInfo{},
		&domain.UserReportInfo{},
	)

	return db, dbErr
}
