package handler

import (
	"net/http"
	"strconv"

	handlerutil "github.com/Nishad4140/ecommerce_project/pkg/api/handlerUtil"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartUsecase services.CartUsecase
}

func NewCartHandler(cartUsecase services.CartUsecase) *CartHandler {
	return &CartHandler{
		cartUsecase: cartUsecase,
	}
}

// -------------------------- Add-To-Cart --------------------------//

func (cr *CartHandler) AddToCart(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("model_id")
	productId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUsecase.AddToCart(productId, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant add product into cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product added into cart",
		Data:       nil,
		Errors:     nil,
	})
}

// -------------------------- Remove-From-Cart --------------------------//

func (cr *CartHandler) RemoveFromCart(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("model_id")
	productId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUsecase.RemoveFromCart(userId, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant remove from cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "item removed from cart",
		Data:       nil,
		Errors:     nil,
	})
}

// -------------------------- List-Cart --------------------------//

func (cr *CartHandler) ListCart(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	cart, err := cr.cartUsecase.ListCart(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "cart items are",
		Data:       cart,
		Errors:     nil,
	})
}
