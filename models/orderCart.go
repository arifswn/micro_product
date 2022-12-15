package models

import "time"

type OrderCart struct {
	Qty       int  `json:"qty" binding:"required"`
	IDProduct uint `json:"id_product" binding:"required"`
}

type OrderCartResponse struct {
	IDOrder   uint          `json:"id_order"`
	IDUser    uint          `json:"id_user"`
	TglOrder  time.Time     `json:"tgl_order"`
	Qty       int           `json:"qty"`
	Total     int           `json:"total"`
	Status    StatusAllowed `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	IDProduct uint          `json:"id_product"`
	Product   ProductOrder  `json:"product"`
}

func FormatResponseOrderCart(product Product, order Order) OrderCartResponse {
	result := OrderCartResponse{
		IDOrder:   order.IDOrder,
		IDUser:    order.IDUser,
		TglOrder:  order.TglOrder,
		Qty:       order.Qty,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		IDProduct: order.IDProduct,
		Product: ProductOrder{
			IDProduct: product.IDProduct,
			Name:      product.Name,
			Stock:     product.Stock,
			Price:     product.Price,
		},
	}
	return result
}
