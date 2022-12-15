package models

import "time"

type ProductUpdate struct {
	Name  string `json:"product_name" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type ProductUpdateResponse struct {
	IDProduct uint      `json:"id_product"`
	Name      string    `json:"product_name"`
	Stock     int       `json:"stock"`
	Price     int       `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatResponseProductUpdate(product Product) ProductUpdateResponse {
	result := ProductUpdateResponse{
		IDProduct: product.IDProduct,
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		UpdatedAt: product.UpdatedAt,
	}
	return result
}
