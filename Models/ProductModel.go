package Models

type Product struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

func (b *Product) TableName() string {
	return "product"
}
