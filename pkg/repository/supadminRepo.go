package repository

import (
	"fmt"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	"gorm.io/gorm"
)

type supadminDatabase struct {
	DB *gorm.DB
}

func NewSupAdminRepository(DB *gorm.DB) interfaces.SupAdminRepository {
	return &supadminDatabase{DB}
}

// -------------------------- Login --------------------------//

func (c *supadminDatabase) SupAdminLogin(email string) (domain.SupAdmins, error) {
	var supadminData domain.SupAdmins
	err := c.DB.Raw("SELECT * FROM sup_admins WHERE email=?", email).Scan(&supadminData).Error
	return supadminData, err
}

// -------------------------- Block-User --------------------------//

func (c *supadminDatabase) BlockUser(body helper.BlockData, adminId int) error {
	// Start a transaction
	tx := c.DB.Begin()
	//Check if the user is there
	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", body.UserId).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}

	// Execute the first SQL command (UPDATE)
	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = ?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Execute the second SQL command (INSERT)
	if err := tx.Exec("INSERT INTO user_block_infos (users_id, reason_for_blocking, blocked_at, blocked_by,block_until) VALUES (?, ?, NOW(), ?, NOW() + INTERVAL '5 Minutes')", body.UserId, body.Reason, adminId).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	// If all commands were executed successfully, return nil
	return nil

}

// -------------------------- UnBlock-User --------------------------//

func (c *supadminDatabase) UnblockUser(id int) error {
	tx := c.DB.Begin()

	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND is_blocked=true)", id).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")
	}
	if err := tx.Exec("UPDATE users SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE user_infos SET reason_for_blocking=$1,blocked_at=NULL,blocked_by=$2 WHERE users_id=$3"
	if err := tx.Exec(query, "", 0, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (c *supadminDatabase) CreateAdmin(adminData helper.AdminData) (response.UserData, error) {
	var userData response.UserData
	insertQuery := `INSERT INTO admins (name,email,mobile,password,created_at)VALUES($1,$2,$3,$4,NOW()) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, adminData.Name, adminData.Email, adminData.Mobile, adminData.Password).Scan(&userData).Error
	return userData, err
}
