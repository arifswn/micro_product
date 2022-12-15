package services

import (
	"micro_product/database"
	"micro_product/models"
	"time"
)

func ProductInputService(input models.ProductInput) (models.Product, error) {
	product := models.Product{}
	db := database.GetDB()

	product.Name = input.Name
	product.Stock = input.Stock
	product.Price = input.Price
	product.CreatedAt = time.Now()

	err := db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func ProductUpdateService(input models.ProductUpdate, idProduct uint) (models.Product, error) {
	product := models.Product{}
	db := database.GetDB()

	err := db.Where("id_product = ?", idProduct).Find(&product).Error
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.UpdatedAt = time.Now()

	err = db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func ProductDeleteService(idProduct uint) (string, error) {
	var product models.Product
	db := database.GetDB()

	err := db.Where("id_product = ?", idProduct).Find(&product).Error
	if err != nil {
		return "Data product tidak ditemukan", err
	}

	err = db.Where("id_product = ?", idProduct).Delete(&product).Error
	if err != nil {
		return "Data product gagal di hapus", err
	}

	return "Data product berhasil dihapus", nil
}

func ProductGetService() ([]models.Product, error) {
	var product []models.Product
	db := database.GetDB()

	err := db.Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func ProductGetServiceById(idProduct uint) (models.Product, error) {
	var product models.Product
	db := database.GetDB()
	err := db.Where("id_product = ?", idProduct).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
