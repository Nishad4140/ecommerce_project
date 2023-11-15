package response

import "time"

type SellerData struct {
	Id     int
	Name   string
	Email  string
	Mobile string
}

type DashBoard struct {
	TotalRevenue        int
	TotalOrders         int
	TotalProductsSelled int
	TotalUsers          int
}

type SalesReport struct {
	Name        string
	PaymentType string
	OrderDate   time.Time
	OrderTotal  int
}
