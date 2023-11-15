package interfaces

import "github.com/Nishad4140/ecommerce_project/pkg/domain"

type PaymentRepository interface {
	ViewPaymentDetails(orderID int) (domain.PaymentDetails, error)
	UpdatePaymentDetails(orderID int, paymentRef string) (domain.PaymentDetails, error)
}
