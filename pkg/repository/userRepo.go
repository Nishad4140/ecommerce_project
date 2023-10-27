package repository

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

//-------------------------- Sign-Up --------------------------//

func (c *userDatabase) UserSignUp(user helper.UserReq) (response.UserData, error) {
	var userData response.UserData
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}

//-------------------------- Login --------------------------//

func (c *userDatabase) UserLogin(email string) (domain.Users, error) {
	var userData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&userData).Error
	return userData, err
}

// -------------------------- View-Profile --------------------------//

func (c *userDatabase) ViewProfile(userID int) (response.UserData, error) {
	var userData response.UserData
	findProfile := `SELECT name,email,mobile FROM users WHERE id=?`
	err := c.DB.Raw(findProfile, userData).Scan(&userData).Error
	return userData, err
}

// -------------------------- Edit-Profile --------------------------//

func (c *userDatabase) EditProfile(userID int, updatingDetails helper.UserReq) (response.UserData, error) {
	var userData response.UserData
	updateProfile := `UPDATE users SET name=$1,email=$2,mobile=$3 WHERE id=$4 RETURNING name,email,mobile`
	err := c.DB.Raw(updateProfile, updatingDetails.Name, updatingDetails.Email, updatingDetails.Mobile, userID).Scan(&userData).Error
	return userData, err
}

// -------------------------- Update-Password --------------------------//

// Find-Old-Password
func (c *userDatabase) FindPassword(userID int) (string, error) {
	var oldPassword string
	err := c.DB.Raw("SELECT password FROM users WHERE id=?", userID).Scan(&oldPassword).Error
	return oldPassword, err
}

// Add-New-Password
func (c *userDatabase) UpdatePassword(userID int, newPassword string) error {
	updatePassword := `UPDATE users SET password=$1 WHERE id=$2`
	err := c.DB.Exec(updatePassword, newPassword, userID).Error
	return err
}

// -------------------------- Add-Address --------------------------//

func (c *userDatabase) AddAddress(userID int, address helper.Address) error {

	if address.IsDefault {
		changeDefault := `UPDATE addresses SET is_default = $1 WHERE users_id=$2 AND is_default=$3`
		err := c.DB.Exec(changeDefault, false, userID, true).Error

		if err != nil {
			return err
		}
	}

	addAddress := `INSERT INTO addresses (users_id,house_number,street,city, district,landmark,pincode,is_default)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	err := c.DB.Exec(addAddress, userID, address.
		House_number,
		address.Street,
		address.City,
		address.District,
		address.Landmark,
		address.Pincode,
		address.IsDefault).Error
	return err
}

// -------------------------- Update-Address --------------------------//

func (c *userDatabase) UpdateAddress(id, addressId int, address helper.Address) error {

	if address.IsDefault {
		changeDefault := `UPDATE addresses SET is_default = $1 WHERE users_id=$2 AND is_default=$3`
		err := c.DB.Exec(changeDefault, false, id, true).Error

		if err != nil {
			return err
		}
	}
	//Update the address
	updateAddress := `UPDATE addresses SET 
		house_number=$1,street=$2,city=$3, district=$4,landmark=$5,pincode=$6,is_default=$7 WHERE users_id=$8 AND id=$9`
	err := c.DB.Exec(updateAddress,
		address.House_number,
		address.Street,
		address.City,
		address.District,
		address.Landmark,
		address.Pincode,
		address.IsDefault,
		id,
		addressId).Error
	return err
}
