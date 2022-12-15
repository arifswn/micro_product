package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type User struct {
	IDUser    uint       `gorm:"primary_key;auto_increment:true"`
	Username  string     `gorm:"not null;uniqueIndex" valid:"required"`
	Email     string     `gorm:"not null;unique" valid:"required, email"`
	Password  string     `gorm:"not null" valid:"required"`
	Age       uint       `gorm:"not null" form:"age" valid:"required"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	//cek umur
	if u.Age < 8 {
		err = errors.New("Make sure you are over 8 years old")
		return
	}

	err = nil
	return
}
