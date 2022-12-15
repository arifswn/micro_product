package models

import "time"

type ProductGetResponse struct {
	IDProduct uint       `json:"id_product"`
	Name      string     `json:"product_name"`
	Stock     int        `json:"stock"`
	Price     int        `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func FormatResponseProduct(product Product) ProductGetResponse {
	result := ProductGetResponse{
		IDProduct: product.IDProduct,
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		DeletedAt: product.DeletedAt,
	}
	return result
}
func FormatResponseProductGet(product []Product) []ProductGetResponse {
	productFormatter := []ProductGetResponse{}
	for _, product := range product {
		responseProductGet := FormatResponseProduct(product)
		productFormatter = append(productFormatter, responseProductGet)
	}
	return productFormatter
}

