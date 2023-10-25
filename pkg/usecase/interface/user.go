package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(user helper.LoginReq) (string, error)
}
