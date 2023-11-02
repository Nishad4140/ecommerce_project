package interfaces

import "github.com/Nishad4140/ecommerce_project/pkg/common/response"

type CartUsecase interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(userId, productId int) error
	ListCart(userId int) (response.ViewCart, error)
}
