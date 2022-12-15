package models

import "time"

type OrderCancel struct {
	IDOrder uint `json:"id_order"`
}

type OrderCancelResponse struct {
	IDOrder     uint          `json:"id_order"`
	IDUser      uint          `json:"id_user"`
	TglOrder    time.Time     `json:"tgl_order"`
	TglCheckout time.Time     `json:"tgl_checkout"`
	TglCancel   time.Time     `json:"tgl_cancel"`
	Qty         int           `json:"qty"`
	Total       int           `json:"total"`
	Status      StatusAllowed `json:"status"`
	UpdatedAt   time.Time     `json:"updated_at"`
	IDProduct   uint          `json:"id_product"`
	Product     ProductOrder  `json:"product"`
}

func FormatResponseOrderCancel(product Product, order Order) OrderCancelResponse {
	result := OrderCancelResponse{
		IDOrder:     order.IDOrder,
		IDUser:      order.IDUser,
		TglOrder:    order.TglOrder,
		TglCheckout: order.TglCheckout,
		TglCancel:   order.TglCancel,
		Qty:         order.Qty,
		Total:       order.Total,
		Status:      order.Status,
		UpdatedAt:   order.UpdatedAt,
		IDProduct:   order.IDProduct,
		Product: ProductOrder{
			IDProduct: product.IDProduct,
			Name:      product.Name,
			Stock:     product.Stock,
			Price:     product.Price,
		},
	}
	return result
}
