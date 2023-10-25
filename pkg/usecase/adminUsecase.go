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

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepo: repo,
	}
}

// -------------------------- Login --------------------------//

func (c *adminUseCase) AdminLogin(admin helper.LoginReq) (string, error) {
	adminData, err := c.adminRepo.AdminLogin(admin.Email)
	if err != nil {
		return "", err
	}

	if admin.Email == "" {
		return "", fmt.Errorf("admin is not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(admin.Password)); err != nil {
		return "", err
	}

	if adminData.IsBlocked {
		return "", fmt.Errorf("admin is blocked")
	}

	claims := jwt.MapClaims{
		"id":  adminData.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// -------------------------- Report-User --------------------------//

func (c *adminUseCase) ReportUser(reason helper.ReportData, adminId int) error {
	err := c.adminRepo.ReportUser(reason, adminId)
	return err
}

// -------------------------- Show-Single-User --------------------------//

func (c *adminUseCase) ShowUser(userID int) (response.UserDetails, error) {
	userData, err := c.adminRepo.ShowUser(userID)
	return userData, err
}

// -------------------------- Show-All-Users --------------------------//

func (c *adminUseCase) ShowAllUser() ([]response.UserDetails, error) {
	userDatas, err := c.adminRepo.ShowAllUser()
	return userDatas, err
}
