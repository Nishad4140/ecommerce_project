package repository

import (
	"fmt"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

//-------------------------- Login --------------------------//

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}

//-------------------------- Report-User --------------------------//

func (c *adminDatabase) ReportUser(reason helper.ReportData, adminId int) error {
	tx := c.DB.Begin()

	var Exists bool
	if err := tx.Raw("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", reason.UserId).Scan(&Exists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !Exists {
		tx.Rollback()
		return fmt.Errorf("no such a user")
	}
	if err := tx.Exec("UPDATE users SET report_count=report_count+1 WHERE id = ?", reason.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("INSERT INTO user_report_infos (users_id, reason_for_reporting, reported_at, reported_by) VALUES (?, ?, NOW(), ?)", reason.UserId, reason.Reason, adminId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

//-------------------------- Show-Single-User --------------------------//

func (c *adminDatabase) ShowUser(userID int) (response.UserDetails, error) {
	var userData response.UserDetails
	qury := `SELECT users.name,
			 users.email, 
			 users.mobile,
			 users.report_count,  
			 users.is_blocked, 
			 block_infos.blocked_by,
			 block_infos.blocked_at,
			 block_infos.reason_for_blocking 
			 FROM users as users 
			 FULL OUTER JOIN user_report_infos as report_infos ON users.id = report_infos.users_id
			 FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id
			 WHERE users.id = $1;`

	err := c.DB.Raw(qury, userID).Scan(&userData).Error
	if err != nil {
		return response.UserDetails{}, err
	}
	if userData.Email == "" {
		return response.UserDetails{}, fmt.Errorf("no such user")
	}
	return userData, nil
}
