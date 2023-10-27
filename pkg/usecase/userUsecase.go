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

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

//-------------------------- Sign-Up --------------------------//

func (c *userUseCase) UserSignUp(user helper.UserReq) (response.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(user)
	return userData, err
}

//-------------------------- Login --------------------------//

func (c *userUseCase) UserLogin(user helper.LoginReq) (string, error) {
	userData, err := c.userRepo.UserLogin(user.Email)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	if userData.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}

	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

//-------------------------- View-Profile --------------------------//

func (c *userUseCase) ViewProfile(userID int) (response.UserData, error) {
	profile, err := c.userRepo.ViewProfile(userID)
	return profile, err
}

//-------------------------- Edit-Profile --------------------------//

func (c *userUseCase) EditProfile(userID int, updatingDetails helper.UserReq) (response.UserData, error) {
	updatedProfile, err := c.userRepo.EditProfile(userID, updatingDetails)
	return updatedProfile, err
}

//-------------------------- Update-Password --------------------------//

func (c *userUseCase) UpdatePassword(userID int, Passwords helper.UpdatePassword) error {

	orginalPassword, err := c.userRepo.FindPassword(userID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(orginalPassword), []byte(Passwords.OldPassword))
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(Passwords.NewPasswoed), 10)
	if err != nil {
		return err
	}
	newPassword := string(hash)

	err = c.userRepo.UpdatePassword(userID, newPassword)
	return err
}

//-------------------------- Add-Address --------------------------//

func (c *userUseCase) AddAddress(userID int, address helper.Address) error {
	err := c.userRepo.AddAddress(userID, address)
	return err
}

//-------------------------- Update-Address --------------------------//

func (c *userUseCase) UpdateAddress(id, addressId int, address helper.Address) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}
