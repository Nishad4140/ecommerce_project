package handler

import (
	"net/http"
	"strconv"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase services.ProductUsecase
}

func NewProductHandler(productUsecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

// -------------------------- Create-Category --------------------------//

func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category helper.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	NewCategoery, err := cr.productUsecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't creat category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Created",
		Data:       NewCategoery,
		Errors:     nil,
	})
}

// -------------------------- Update-Category --------------------------//

func (cr *ProductHandler) UpdatCategory(c *gin.Context) {
	var category helper.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedCategory, err := cr.productUsecase.UpdatCategory(category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Updated",
		Data:       updatedCategory,
		Errors:     nil,
	})
}

// -------------------------- Delete-Category --------------------------//

func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	parmasId := c.Param("category_id")
	id, err := strconv.Atoi(parmasId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.productUsecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't dlete category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category deleted",
		Data:       nil,
		Errors:     nil,
	})

}

// -------------------------- List-All-Category --------------------------//

func (cr *ProductHandler) ListAllCategories(c *gin.Context) {
	categories, err := cr.productUsecase.ListAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Ctegories are",
		Data:       categories,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Category --------------------------//

func (cr *ProductHandler) ListCategory(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	category, err := cr.productUsecase.ListCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category is",
		Data:       category,
		Errors:     nil,
	})
}

// -------------------------- Create-Brand --------------------------//

func (cr *ProductHandler) AddProduct(c *gin.Context) {
	var product helper.Brand
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newProduct, err := cr.productUsecase.AddProduct(product)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't add product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Added",
		Data:       newProduct,
		Errors:     nil,
	})

}

// -------------------------- Update-Brand --------------------------//

func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var product helper.Brand
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedProduct, err := cr.productUsecase.UpdateProduct(id, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant update brand",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, response.Response{
		StatusCode: 200,
		Message:    "Brand updated",
		Data:       updatedProduct,
		Errors:     nil,
	})

}

// -------------------------- Delete-Brand --------------------------//

func (cr *ProductHandler) DeleteProduct(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productUsecase.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// -------------------------- List-All-Brand --------------------------//

func (cr *ProductHandler) ListAllProduct(c *gin.Context) {

	products, err := cr.productUsecase.ListAllProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       products,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Brand --------------------------//

func (cr *ProductHandler) ListProduct(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	product, err := cr.productUsecase.ListProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       product,
		Errors:     nil,
	})
}

// -------------------------- Add-Model --------------------------//

func (cr *ProductHandler) AddModel(c *gin.Context) {
	var model helper.Model
	err := c.Bind(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newModel, err := cr.productUsecase.AddModel(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant create",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product created",
		Data:       newModel,
		Errors:     nil,
	})
}

// -------------------------- Update-Model --------------------------//

func (cr *ProductHandler) UpdateModel(c *gin.Context) {
	var model helper.Model
	err := c.Bind(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedItem, err := cr.productUsecase.UpdateModel(id, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update productitem",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "productitem updated",
		Data:       updatedItem,
		Errors:     nil,
	})
}

// -------------------------- Delete-Model --------------------------//

func (cr *ProductHandler) DeleteModel(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productUsecase.DeleteModel(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "item deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// -------------------------- List-All-Model --------------------------//

func (cr *ProductHandler) ListAllModel(c *gin.Context) {

	productItems, err := cr.productUsecase.ListAllModel()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't disaply items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product items are",
		Data:       productItems,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Model --------------------------//

func (cr *ProductHandler) ListModel(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	productItem, err := cr.productUsecase.ListModel(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       productItem,
		Errors:     nil,
	})

}

// -------------------------- Upload-Image --------------------------//

func (cr *ProductHandler) UploadImage(c *gin.Context) {

	id := c.Param("id")
	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	// Multipart form
	form, _ := c.MultipartForm()

	files := form.File["images"]

	for _, file := range files {
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "asset/uploads/"+file.Filename)

		err := cr.productUsecase.UploadImage(file.Filename, productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "cant upload images",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}
}
