package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
)

type SupAdminUseCase interface {
	SupAdminLogin(supadmin helper.LoginReq) (string, error)
	BlockUser(body helper.BlockData, adminId int) error
	UnblockUser(id int) error
}
