package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type Product struct {
	IDProduct uint       `gorm:"primary_key;auto_increment:true"`
	Name      string     `gorm:"not null;uniqueIndex" valid:"required"`
	Stock     int        `gorm:"not null" valid:"required"`
	Price     int        `gorm:"not null" valid:"required"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	// inisialisasi struct product
	var product Product
	// find product name
	errInputProduct := tx.Model(&Product{}).Where("name = ?", p.Name).First(&product).Error
	// cek product
	if errInputProduct == nil {
		return errors.New("Product name already exists")
	}

	err = nil
	return
}
