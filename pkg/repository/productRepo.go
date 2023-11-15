package repository

import (
	"fmt"
	"strings"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &ProductDatabase{DB}
}

// -------------------------- Create-Category --------------------------//

func (c *ProductDatabase) CreateCategory(category helper.Category) (response.Category, error) {
	var newCategory response.Category
	query := `INSERT INTO categories (category_name,created_at)VAlues($1,NOW())RETURNING id,category_name`
	err := c.DB.Raw(query, category.Name).Scan(&newCategory).Error
	return newCategory, err
}

// -------------------------- Update-Category --------------------------//

func (c *ProductDatabase) UpdatCategory(category helper.Category, id int) (response.Category, error) {
	var updatedCategory response.Category
	query := `UPDATE categories
	          SET category_name = $1, updated_at = NOW()
	          WHERE id = $2
	          RETURNING id, category_name`
	err := c.DB.Raw(query, category.Name, id).Scan(&updatedCategory).Error
	if err != nil {
		return response.Category{}, err
	}
	if updatedCategory.Id == 0 {
		return response.Category{}, fmt.Errorf("no such category to update")
	}
	return updatedCategory, nil
}

// -------------------------- Delete-Category --------------------------//

func (c *ProductDatabase) DeleteCategory(id int) error {
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return fmt.Errorf("no category found")
	}
	query := `DELETE FROM categories WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

// -------------------------- List-All-Category --------------------------//

func (c *ProductDatabase) ListAllCategories() ([]response.Category, error) {
	var categories []response.Category
	query := `SELECT * FROM categories`
	err := c.DB.Raw(query).Scan(&categories).Error
	return categories, err
}

// -------------------------- List-Single-Category --------------------------//

func (c *ProductDatabase) ListCategory(id int) (response.Category, error) {
	var category response.Category
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, id).Scan(&exits)
	if !exits {
		return response.Category{}, fmt.Errorf("no category found")
	}
	query := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	if err != nil {
		return response.Category{}, err
	}
	if category.Id == 0 {
		return response.Category{}, fmt.Errorf("no such category")
	}
	return category, nil
}

// -------------------------- Create-Product --------------------------//

func (c *ProductDatabase) AddProduct(product helper.Brand) (response.Brand, error) {
	var newProduct response.Brand
	var exits bool

	query1 := `select exists(select 1 from categories where id=?)`
	c.DB.Raw(query1, product.CategoryId).Scan(&exits)
	if !exits {
		return response.Brand{}, fmt.Errorf("no category found")
	}

	query := `INSERT INTO brands (brand_name,description,category_id,created_at)
		VALUES ($1,$2,$3,NOW())
		RETURNING id,brand_name AS name,description,category_id`
	err := c.DB.Raw(query, product.Name, product.Description, product.CategoryId).
		Scan(&newProduct).Error
	return newProduct, err
}

// -------------------------- Update-Product --------------------------//

func (c *ProductDatabase) UpdateProduct(id int, product helper.Brand) (response.Brand, error) {
	var updatedProduct response.Brand
	query2 := `UPDATE brands SET brand_name=$1,description=$2,category_id=$3,updated_at=NOW() WHERE id=$4
		RETURNING id,brand_name,description,category_id`
	err := c.DB.Raw(query2, product.Name, product.Description, product.CategoryId, id).
		Scan(&updatedProduct).Error
	if err != nil {
		return response.Brand{}, err
	}
	if updatedProduct.Id == 0 {
		return response.Brand{}, fmt.Errorf("there is no such product")
	}
	return updatedProduct, nil
}

// -------------------------- Delete-Product --------------------------//

func (c *ProductDatabase) DeleteProduct(id int) error {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM brands WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("there is no such product to delete")
	}
	query := `DELETE FROM brands WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}

// -------------------------- List-All-Product --------------------------//

func (c *ProductDatabase) ListAllProduct() ([]response.Brand, error) {
	var products []response.Brand
	getProductDetails := `SELECT p.id,p.brand_name AS name,
		p.description,
		c.category_name
		 FROM brands p JOIN categories c ON p.category_id=c.id `

	err := c.DB.Raw(getProductDetails).Scan(&products).Error
	if err != nil {
		return []response.Brand{}, err
	}
	return products, nil
}

// -------------------------- List-Single-Product --------------------------//

func (c *ProductDatabase) ListProduct(id int) (response.Brand, error) {
	var product response.Brand
	query := `SELECT p.id,p.brand_name AS name,p.description,c.category_name FROM brands p 
		JOIN categories c ON p.category_id=c.id WHERE p.id=$1`
	err := c.DB.Raw(query, id).Scan(&product).Error
	if err != nil {
		return response.Brand{}, err
	}
	if product.Id == 0 {
		return response.Brand{}, fmt.Errorf("there is no such product")
	}
	return product, err
}

// -------------------------- Add-Model --------------------------//

func (c *ProductDatabase) AddModel(model helper.Model) (response.Model, error) {
	var newModel response.Model
	query := `INSERT INTO models (brand_id,
		model_name,
		sku,
		qty_in_stock,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price,
		image,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,NOW())
		RETURNING 
		id,
		model_name,
		brand_id,
		sku,
		qty_in_stock,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price,
		image`
	err := c.DB.Raw(query, model.Brand_id,
		model.Model_name,
		model.Sku,
		model.Qty,
		model.Color,
		model.Ram,
		model.Battery,
		model.Screen_size,
		model.Storage,
		model.Camera,
		model.Price,
		model.Image).Scan(&newModel).Error
	if err != nil {
		return newModel, err
	}
	selectQuery := `SELECT p.brand_name,
		p.description,
		c.category_name, 
		pi.*
		FROM brands p 
		JOIN categories c ON p.category_id=c.id 
		JOIN models pi ON p.id=pi.brand_id 
		WHERE pi.id=$1`
	err = c.DB.Raw(selectQuery, newModel.Id).Scan(&newModel).Error
	return newModel, err

}

// -------------------------- Update-Model --------------------------//

func (c *ProductDatabase) UpdateModel(id int, model helper.Model) (response.Model, error) {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM models WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return response.Model{}, fmt.Errorf("there is no such model to update")
	}

	var updatedModel response.Model
	query := `UPDATE models SET 
	brand_id=$1,
	model_name=$13, -- Corrected placeholder
	sku=$2,
	qty_in_stock=$3,
	image=$4,
	color=$5,
	ram=$6,
	battery=$7,
	screen_size=$8,
	storage=$9,
	camera=$10,
	price=$11,
	updated_at=NOW()
	WHERE id=$12
	RETURNING
		id,
		model_name,
		brand_id,
		sku,
		qty_in_stock,
		color,
		ram,
		battery,
		screen_size,
		storage,
		camera,
		price,
		image`

	err := c.DB.Raw(query,
		model.Brand_id,
		model.Sku,
		model.Qty,
		model.Image,
		model.Color,
		model.Ram,
		model.Battery,
		model.Screen_size,
		model.Storage,
		model.Camera,
		model.Price,
		id,
		model.Model_name).Scan(&updatedModel).Error

	return updatedModel, err
}

// -------------------------- Delete-Model --------------------------//

func (c *ProductDatabase) DeleteModel(id int) error {
	var exists bool
	isExists := `SELECT EXISTS (SELECT 1 FROM models WHERE id=$1)`
	c.DB.Raw(isExists, id).Scan(&exists)
	if !exists {
		return fmt.Errorf("there is no such model to delete")
	}
	query := `DELETE FROM models WHERE id=?`
	err := c.DB.Exec(query, id).Error
	return err
}

// -------------------------- List-All-Model --------------------------//

func (c *ProductDatabase) ListAllModel(queryParams helper.QueryParams) ([]response.Model, error) {
	var models []response.Model
	getProductItemDetails := `SELECT p.brand_name,
		p.description,
		c.category_name, 
		pi.*
		FROM brands p 
		JOIN categories c ON p.category_id=c.id 
		JOIN models pi ON pi.brand_id=p.id`

	if queryParams.Query != "" && queryParams.Filter != "" {
		getProductItemDetails = fmt.Sprintf("%s WHERE LOWER(%s) LIKE '%%%s%%'", getProductItemDetails, queryParams.Filter, strings.ToLower(queryParams.Query))
	}

	if queryParams.SortBy != "" {
		if queryParams.SortDesc {
			getProductItemDetails = fmt.Sprintf("%s ORDER BY %s DESC", getProductItemDetails, queryParams.SortBy)
		} else {
			getProductItemDetails = fmt.Sprintf("%s ORDER BY %s ASC", getProductItemDetails, queryParams.SortBy)
		}
	} else {
		getProductItemDetails = fmt.Sprintf("%s ORDER BY p.created_at DESC", getProductItemDetails)
	}
	//to set the page number and the qty that need to display in a single responce
	if queryParams.Limit != 0 && queryParams.Page != 0 {
		getProductItemDetails = fmt.Sprintf("%s LIMIT %d OFFSET %d", getProductItemDetails, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
	}
	if queryParams.Limit == 0 || queryParams.Page == 0 {
		getProductItemDetails = fmt.Sprintf("%s LIMIT 10 OFFSET 0", getProductItemDetails)
	}

	err := c.DB.Raw(getProductItemDetails).Scan(&models).Error
	return models, err
}

// -------------------------- List-Single-Model --------------------------//

func (c *ProductDatabase) ListModel(id int) (response.Model, error) {
	var productItem response.Model
	query := `SELECT p.brand_name,
	p.description,
	p.brand_name,
	c.category_name, 
	pi.*
	FROM brands p 
	JOIN categories c ON p.category_id=c.id 
	JOIN models pi ON p.id=pi.brand_id 
	WHERE pi.id=$1`
	err := c.DB.Raw(query, id).Scan(&productItem).Error
	if err != nil {
		return response.Model{}, err
	}
	if productItem.Id == 0 {
		return response.Model{}, fmt.Errorf("there is no such product item")
	}
	getImages := `SELECT file_name FROM images WHERE model_id=$1`
	err = c.DB.Raw(getImages, id).Scan(&productItem.Image).Error
	if err != nil {
		return response.Model{}, err
	}
	return productItem, nil
}

// -------------------------- Uploaded-Model --------------------------//

func (c *ProductDatabase) UploadImage(filepath string, productId int) error {
	uploadImage := `INSERT INTO images (model_id,file_name)VALUES($1,$2)`
	err := c.DB.Exec(uploadImage, productId, filepath).Error
	return err
}
