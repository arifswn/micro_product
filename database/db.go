package database

import (
	"fmt"
	"micro_product/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Can't load configuration : " + err.Error())
	}

	// load db configuration
	DBName := os.Getenv("DB_NAME")
	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")

	// conn db
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", DBUsername, DBPassword, DBHost, DBName))
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
