package interfaces

import (
	helper "github.com/Nishad4140/ecommerce_project/pkg/common/helperStruct"
	"github.com/Nishad4140/ecommerce_project/pkg/domain"
)

type PaymentUseCase interface {
	CreateRazorpayPayment(userId, orderId int) (domain.Orders, string, error)
	UpdatePaymentDetails(paymentVerifier helper.PaymentVerification) error
}
