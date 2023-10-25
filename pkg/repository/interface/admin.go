package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
)

type AdminRepository interface {
	AdminLogin(email string) (domain.Admins, error)
	ReportUser(reason helper.ReportData, adminId int) error
	ShowUser(userID int) (response.UserDetails, error)
}
