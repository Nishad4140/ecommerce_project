package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nishad4140/ecommerce_project/pkg/common/response"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCase
}

func NewPaymentHandler(paymentUseCase services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

func (cr *PaymentHandler) CreateRazorpayPayment(c *gin.Context) {
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find order id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	fmt.Println(paramsId)
	// userId, err := handlerUtil.GetUserIdFromContext(c)

	// fmt.Println("1", userId)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, response.Response{
	// 		StatusCode: 400,
	// 		Message:    "Can't find UserId",
	// 		Data:       nil,
	// 		Errors:     err.Error(),
	// 	})
	// 	return
	// }
	userId := 14
	order, razorpayID, err := cr.paymentUseCase.CreateRazorpayPayment(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't complete order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.HTML(200, "app.html", gin.H{
		"UserID":       userId,
		"total_price":  order.OrderTotal,
		"total":        order.OrderTotal,
		"orderData":    order.Id,
		"orderid":      razorpayID,
		"amount":       order.OrderTotal,
		"Email":        "nishadshanid40@gmail.com",
		"Phone_Number": "8848994140",
	})
}
