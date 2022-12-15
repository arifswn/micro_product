package models

import "time"

type ProductInput struct {
	Name  string `json:"product_name" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type ProductInputResponse struct {
	IDProduct uint      `json:"id_product"`
	Name      string    `json:"product_name"`
	Stock     int       `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatResponseProductInput(product Product) ProductInputResponse {
	result := ProductInputResponse{
		IDProduct: product.IDProduct,
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
	}
	return result
}
