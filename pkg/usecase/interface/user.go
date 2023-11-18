package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(user helper.LoginReq) (string, error)
	ForgotPassword(forgotPass helper.ForgotPassword) error
	ViewProfile(userID int) (response.Userprofile, error)
	EditProfile(userID int, updatingDetails helper.UpdateProfile) (response.Userprofile, error)
	UpdatePassword(userID int, Passwords helper.UpdatePassword) error
	AddAddress(id int, address helper.Address) error
	UpdateAddress(id, addressId int, address helper.Address) error
	CreateWallet(id int) error
	VerifyWallet(id int) error
}
