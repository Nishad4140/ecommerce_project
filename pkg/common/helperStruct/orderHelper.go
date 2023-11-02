package helper

type Cart struct {
	Id     int
	Total int
}

type CartItems struct {
	ModelId    int
	Quantity   int
	Price      int
	QtyInStock int
}

type UpdateOrder struct {
	OrderId       uint
	OrderStatusID uint
}
