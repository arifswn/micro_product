package models

import (
	"time"

	dtoMapper "github.com/dranikpg/dto-mapper"
)

type OrderGet struct {
	IDOrder uint `json:"id_order"`
}

type OrderGetResponse struct {
	IDOrder     uint          `json:"id_order"`
	IDUser      uint          `json:"id_user"`
	TglOrder    time.Time     `json:"tgl_order"`
	TglCheckout time.Time     `json:"tgl_checkout"`
	TglCancel   time.Time     `json:"tgl_cancel"`
	Qty         int           `json:"qty"`
	Total       int           `json:"total"`
	Status      StatusAllowed `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	IDProduct   uint          `json:"id_product"`
	Product     ProductOrder  `json:"product"`
}

func FormatResponseOrder(order Order) OrderGetResponse {
	result := OrderGetResponse{
		IDOrder:     order.IDOrder,
		IDUser:      order.IDUser,
		TglOrder:    order.TglOrder,
		TglCheckout: order.TglCheckout,
		TglCancel:   order.TglCancel,
		Qty:         order.Qty,
		Total:       order.Total,
		Status:      order.Status,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		IDProduct:   order.IDProduct,
		Product: ProductOrder{
			IDProduct: order.Product.IDProduct,
			Name:      order.Product.Name,
			Stock:     order.Product.Stock,
			Price:     order.Product.Price,
		},
	}
	return result
}

func FormatResponseOrderAll(order []Order) []OrderGetResponse {
	var orderGetResponse []OrderGetResponse
	dtoMapper.Map(&orderGetResponse, order)
	return orderGetResponse
}
