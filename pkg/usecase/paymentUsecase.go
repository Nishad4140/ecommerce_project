package usecase

import (
	"fmt"

	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
	interfaces "github.com/Nishad4140/ecommerce_project/pkg/repository/interface"
	services "github.com/Nishad4140/ecommerce_project/pkg/usecase/interface"
	"github.com/razorpay/razorpay-go"
)

type PaymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
	orderRepo   interfaces.OrderRepository
	cfg         config.Config
}

func NewPaymentuseCase(paymentRepo interfaces.PaymentRepository, orderRepo interfaces.OrderRepository, cfg config.Config) services.PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		cfg:         cfg,
	}
}

func (c *PaymentUseCase) CreateRazorpayPayment(orderId int) (domain.Orders, string, int, error) {
	paymentDetails, err := c.paymentRepo.ViewPaymentDetails(orderId)
	if err != nil {
		return domain.Orders{}, "", 0, err
	}

	if paymentDetails.PaymentStatusID == 3 {
		return domain.Orders{}, "", 0, fmt.Errorf("payment already completed")
	}
	userId, err := c.orderRepo.UserIdFromOrder(orderId)
	if err != nil {
		return domain.Orders{}, "", 0, err
	}
	fmt.Println("user id ", userId)
	//fetch order details from the db
	order, err := c.orderRepo.ListOrder(userId, orderId)
	if err != nil {
		return domain.Orders{}, "", userId, err
	}
	fmt.Println(order.Id)
	if order.Id == 0 {
		return domain.Orders{}, "", userId, fmt.Errorf("no such order found")
	}
	client := razorpay.NewClient(c.cfg.RAZORPAYID, c.cfg.RAZORPAYSECRET)

	data := map[string]interface{}{
		"amount":   order.OrderTotal * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Orders{}, "", userId, err
	}

	value := body["id"]
	razorpayID := value.(string)
	return order, razorpayID, userId, err
}

func (c *PaymentUseCase) UpdatePaymentDetails(paymentVerifier helper.PaymentVerification) error {
	paymentDetails, err := c.paymentRepo.ViewPaymentDetails(paymentVerifier.OrderID)
	if err != nil {
		return err
	}
	if paymentDetails.ID == 0 {
		return fmt.Errorf("no order found")
	}

	if paymentDetails.OrderTotal != paymentVerifier.Total {
		return fmt.Errorf("payment amount and order amount does not match")
	}
	updatedPayment, err := c.paymentRepo.UpdatePaymentDetails(paymentVerifier.OrderID, paymentVerifier.PaymentRef)
	if err != nil {
		return err
	}
	if updatedPayment.ID == 0 {
		return fmt.Errorf("failed to update payment details")
	}
	return nil
}
