package usecase

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
)

type ProductUsecase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(productRepo interfaces.ProductRepository) services.ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
	}
}

// -------------------------- Create-Category --------------------------//

func (c *ProductUsecase) CreateCategory(category helper.Category) (response.Category, error) {
	newCategory, err := c.productRepo.CreateCategory(category)
	return newCategory, err
}

// -------------------------- Update-Category --------------------------//

func (c *ProductUsecase) UpdatCategory(category helper.Category, id int) (response.Category, error) {
	updatedCategory, err := c.productRepo.UpdatCategory(category, id)
	return updatedCategory, err
}

// -------------------------- Delete-Category --------------------------//

func (c *ProductUsecase) DeleteCategory(id int) error {
	err := c.productRepo.DeleteCategory(id)
	return err
}

// -------------------------- List-All-Category --------------------------//

func (c *ProductUsecase) ListAllCategories() ([]response.Category, error) {
	categories, err := c.productRepo.ListAllCategories()
	return categories, err
}

// -------------------------- List-Single-Category --------------------------//

func (c *ProductUsecase) ListCategory(id int) (response.Category, error) {
	category, err := c.productRepo.ListCategory(id)
	return category, err
}

// -------------------------- Create-Product --------------------------//

func (c *ProductUsecase) AddProduct(product helper.Brand) (response.Brand, error) {
	newProduct, err := c.productRepo.AddProduct(product)
	return newProduct, err
}

// -------------------------- Update-Product --------------------------//

func (c *ProductUsecase) UpdateProduct(id int, product helper.Brand) (response.Brand, error) {
	updatedProduct, err := c.productRepo.UpdateProduct(id, product)
	return updatedProduct, err
}

// -------------------------- Delete-Product --------------------------//

func (c *ProductUsecase) DeleteProduct(id int) error {
	err := c.productRepo.DeleteProduct(id)
	return err
}

// -------------------------- List-All-Product --------------------------//

func (c *ProductUsecase) ListAllProduct() ([]response.Brand, error) {
	products, err := c.productRepo.ListAllProduct()
	return products, err
}

// -------------------------- List-Single-Product --------------------------//

func (c *ProductUsecase) ListProduct(id int) (response.Brand, error) {
	product, err := c.productRepo.ListProduct(id)
	return product, err
}

// -------------------------- Add-Model --------------------------//

func (c *ProductUsecase) AddModel(model helper.Model) (response.Model, error) {
	newModel, err := c.productRepo.AddModel(model)
	return newModel, err
}
