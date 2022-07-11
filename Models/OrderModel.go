package Models

func (b *Order) TableName() string {
	return "order"
}

type Order struct {
	Id         int    `json:"id"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	Quantity   string `json:"quantity"`
	Status     string `json:"status"`
}
