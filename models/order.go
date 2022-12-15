package models

import (
	"database/sql/driver"
	"time"
)

type StatusAllowed string

const (
	cart     StatusAllowed = "cart"
	checkout StatusAllowed = "checkout"
	cancel   StatusAllowed = "cancel"
)

func (st *StatusAllowed) Scan(value interface{}) error {
	*st = StatusAllowed(value.([]byte))
	return nil
}

func (st StatusAllowed) Value() (driver.Value, error) {
	return string(st), nil
}

type Order struct {
	IDOrder     uint          `gorm:"primary_key;auto_increment:true"`
	IDUser      uint          `gorm:"not null"`
	TglOrder    time.Time     `gorm:"not null"`
	TglCheckout time.Time     `gorm:"not null"`
	TglCancel   time.Time     `gorm:"not null"`
	Qty         int           `gorm:"not null" valid:"required"`
	Total       int           `gorm:"not null" valid:"required"`
	Status      StatusAllowed `gorm:"column:status_allowed" sql:"type:ENUM('cart', 'checkout', 'cancel')"`
	CreatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time    `sql:"index"`
	IDProduct   uint          `gorm:"references:id_product"`
	Product     Product       `gorm:"Foreignkey:IDProduct;association_foreignkey:id_product;" json:"id_product"`
}
type ProductOrder struct {
	IDProduct uint   `json:"id_product"`
	Name      string `json:"product_name"`
	Stock     int    `json:"stock"`
	Price     int    `json:"price"`
}
