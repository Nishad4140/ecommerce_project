package usecase

import (
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
)

type CartUseCase struct {
	cartRepo interfaces.CartRepository
}

func NewCartUsecase(cartRepo interfaces.CartRepository) services.CartUsecase {
	return &CartUseCase{
		cartRepo: cartRepo,
	}
}

// -------------------------- Create-Cart --------------------------//

func (c *CartUseCase) CreateCart(id int) error {
	err := c.cartRepo.CreateCart(id)
	return err
}

// -------------------------- Add-To-Cart --------------------------//

func (c *CartUseCase) AddToCart(productId, userId int) error {
	err := c.cartRepo.AddToCart(productId, userId)
	return err
}

// -------------------------- Remove-From-Cart --------------------------//

func (c *CartUseCase) RemoveFromCart(userId, productId int) error {
	err := c.cartRepo.RemoveFromCart(userId, productId)
	return err
}

// -------------------------- List-Cart --------------------------//

func (c *CartUseCase) ListCart(userId int) (response.ViewCart, error) {
	items, err := c.cartRepo.ListCart(userId)
	return items, err
}
