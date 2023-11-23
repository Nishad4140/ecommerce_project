package usecase

import (
	"fmt"
	"time"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type supadminUseCase struct {
	supadminRepo interfaces.SupAdminRepository
}

func NewSupAdminUseCase(repo interfaces.SupAdminRepository) services.SupAdminUseCase {
	return &supadminUseCase{
		supadminRepo: repo,
	}
}

// -------------------------- Login --------------------------//

func (c *supadminUseCase) SupAdminLogin(supadmin helper.LoginReq) (string, error) {
	supadminData, err := c.supadminRepo.SupAdminLogin(supadmin.Email)
	if err != nil {
		return "", err
	}

	if supadmin.Email == "" {
		return "", fmt.Errorf("supadmin is not found")
	}

	if supadmin.Password != supadminData.Password {
		return "", fmt.Errorf("incorrect password")
	}

	// if err = bcrypt.CompareHashAndPassword([]byte(supadminData.Password), []byte(supadmin.Password)); err != nil {
	// 	return "", err
	// }

	claims := jwt.MapClaims{
		"id":  supadminData.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// -------------------------- Block-User --------------------------//

func (c *supadminUseCase) BlockUser(body helper.BlockData, adminId int) error {
	err := c.supadminRepo.BlockUser(body, adminId)
	return err
}

// -------------------------- UnBlock-User --------------------------//

func (c *supadminUseCase) UnblockUser(id int) error {
	err := c.supadminRepo.UnblockUser(id)
	return err
}

func (c *supadminUseCase) CreateAdmin(adminData helper.AdminData) (response.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(adminData.Password), 10)
	if err != nil {
		return response.UserData{}, err
	}
	adminData.Password = string(hash)
	userData, err := c.supadminRepo.CreateAdmin(adminData)
	return userData, err
}
