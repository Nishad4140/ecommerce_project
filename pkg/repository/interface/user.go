package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
)

type UserRepository interface {
	UserSignUp(user helper.UserReq) (response.UserData, error)
	UserLogin(email string) (domain.Users, error)
}
