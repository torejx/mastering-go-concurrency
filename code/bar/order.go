package bar

type Order struct {
	ID     int
	Coffee Coffee
}

func NewOrder(id int, coffee Coffee) Order {
	return Order{
		ID:     id,
		Coffee: coffee,
	}
}

type Orders []Order
