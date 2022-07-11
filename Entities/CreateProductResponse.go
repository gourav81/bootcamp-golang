package Entities

type CreateProductResponse struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Message     string `json:"message"`
}
