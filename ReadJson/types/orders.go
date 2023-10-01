package types

type Order struct {
	OrderId     string  `json:"order_id"`
	CustomerId  string  `json:"customer_id"`
	Items       []Item  `json:"items"`
	TotalCharge float64 `json:"total_charge"`
}

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type OrderList struct {
	Orders []Order `json:"orders"`
}
