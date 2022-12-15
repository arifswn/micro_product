package database

import (
	"fmt"
	"micro_product/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// config db
const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3333"
	dbname   = "db_micro"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {

	// conn db
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", username, password, hostname, dbname))
	if err != nil {
		panic("Failed connect to database : " + err.Error())
	}

	AutoMigrate()
	fmt.Println("Database Connected...")
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate() {
	// auto migrate models to database table
	err := db.Debug().AutoMigrate(models.User{}, models.Product{}, models.Order{}).Error
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to migrate models : %s", err.Error())
		panic(errorMessage)
	}
	return
}
