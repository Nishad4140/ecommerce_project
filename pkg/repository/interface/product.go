package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
)

type ProductRepository interface {
	CreateCategory(category helper.Category) (response.Category, error)
	UpdatCategory(category helper.Category, id int) (response.Category, error)
	DeleteCategory(id int) error
	ListAllCategories() ([]response.Category, error)
	ListCategory(id int) (response.Category, error)

	AddProduct(product helper.Brand) (response.Brand, error)
	UpdateProduct(id int, product helper.Brand) (response.Brand, error)
	DeleteProduct(id int) error
	ListAllProduct() ([]response.Brand, error)
	ListProduct(id int) (response.Brand, error)

	AddModel(model helper.Model) (response.Model, error)
}
