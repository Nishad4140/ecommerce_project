package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
)

type OrderUseCase interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListOrder(userId, orderId int) (domain.Orders, error)
	ListAllOrders(userId int) ([]domain.Orders, error)
	ReturnOrder(userId, orderId int) (int, error)
	UpdateOrder(UpdateOrder helper.UpdateOrder) error
}
