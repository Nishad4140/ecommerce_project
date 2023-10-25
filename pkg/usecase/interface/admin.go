package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
)

type AdminUseCase interface {
	AdminLogin(admin helper.LoginReq) (string, error)
	ReportUser(reason helper.ReportData, adminID int) error
	ShowUser(userID int) (response.UserDetails, error)
}
