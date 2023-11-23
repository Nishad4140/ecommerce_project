package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
)

type SupAdminRepository interface {
	SupAdminLogin(email string) (domain.SupAdmins, error)
	BlockUser(body helper.BlockData, adminId int) error
	UnblockUser(id int) error
	CreateAdmin(adminData helper.AdminData) (response.UserData, error)
}
