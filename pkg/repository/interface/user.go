package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
)

type UserRepository interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(email string) (domain.Users, error)
	UserDetails(email string) (response.UserData, error)
	ViewProfile(userID int) (response.Userprofile, error)
	EditProfile(userID int, updatingDetails helper.UpdateProfile) (response.Userprofile, error)
	FindPassword(id int) (string, error)
	UpdatePassword(id int, newPassword string) error
	AddAddress(id int, address helper.Address) error
	UpdateAddress(userID, addressID int, address helper.Address) error
	CreateWallet(id int) error
	VerifyWallet(id int) error
}
