package helper

type Order struct {
	PatmentType int
}

type Cart struct {
	Id    int
	Total int
}

type CartItems struct {
	ModelId    int
	Quantity   int
	Price      int
	QtyInStock int
}

type OrderItems struct {
	ModelId    int
	Quantity   int
	Price      int
	QtyInStock int
}

type UpdateOrder struct {
	OrderId       uint `json:"orderid"`
	OrderStatusID uint `json:"order_status_id"`
}
