package handler

import (
	"net/http"
	"strconv"

	handlerutil "github.com/Nishad4140/ecommerce_project/pkg/api/handlerUtil"
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(orderUseCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

//-------------------------- Order-All --------------------------//

func (cr *OrderHandler) OrderAll(c *gin.Context) {
	paramsId := c.Param("payment_id")
	paymentTypeId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
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
	order, err := cr.orderUseCase.OrderAll(userId, paymentTypeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})
}

//-------------------------- Cancel-Order --------------------------//

func (cr *OrderHandler) UserCancelOrder(c *gin.Context) {
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
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.orderUseCase.UserCancelOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't cancel order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order canceld",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- List-Order --------------------------//

func (cr *OrderHandler) ListOrder(c *gin.Context) {
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
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.orderUseCase.ListOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       order,
		Errors:     nil,
	})
}

//-------------------------- List-All-Order --------------------------//

func (cr *OrderHandler) ListAllOrders(c *gin.Context) {
	Id, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	orders, err := cr.orderUseCase.ListAllOrders(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       orders,
		Errors:     nil,
	})
}

//-------------------------- Return-Order --------------------------//

func (cr *OrderHandler) ReturnOrder(c *gin.Context) {
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
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	returnAmount, err := cr.orderUseCase.ReturnOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't return order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order returnd ",
		Data:       returnAmount,
		Errors:     nil,
	})
}

//-------------------------- Update-Order --------------------------//

func (cr *OrderHandler) UpdateOrder(c *gin.Context) {
	var UpdateOrder helper.UpdateOrder
	err := c.BindJSON(&UpdateOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	err = cr.orderUseCase.UpdateOrder(UpdateOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order updated ",
		Data:       nil,
		Errors:     nil,
	})
}
