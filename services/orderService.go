package services

import (
	"errors"
	"micro_product/database"
	"micro_product/models"
	"time"
)

func OrderCartService(input models.OrderCart, userId uint) (models.Product, models.Order, error) {
	order := models.Order{}
	product := models.Product{}
	db := database.GetDB()

	idProduct := input.IDProduct

	// cek id product
	err := db.Where("id_product = ?", idProduct).Find(&product).Error
	if err != nil {
		return product, order, errors.New("Product tidak ditemukkan")
	}

	// cek stock product jika sisa 0, tidak bisa melakukan penambahan cart
	if product.Stock == 0 {
		return product, order, errors.New("Stock product sisa 0. Tidak bisa melakukan penambahan cart")
	}

	// inisialisasi status product
	var statusAlready bool = false
	// cek order jika product sudah ada di order, maka akan di update
	err = db.Where("id_product = ? AND id_user = ? AND status_allowed != ? AND status_allowed != ?",
		idProduct,
		userId,
		models.StatusAllowed("checkout"),
		models.StatusAllowed("cancel")).Find(&order).Error
	if err == nil {
		// set status jika product sudah ada di order
		statusAlready = true
	}

	// ambil data price product
	priceProduct := product.Price
	// ambil status order
	statusOrder := order.Status
	// set status default untuk update data order
	var statusExecute bool = false

	// cek kondisi status jika product sudah ada di order
	if !statusAlready {
		// buat data order baru
		statusExecute = true
	} else {
		// cek kondisi status order jika statusnya cancel & checkout
		if statusOrder == "cancel" || statusOrder == "checkout" {
			// buat data order baru
			statusExecute = true
		}
	}

	// cek kondis jika status data sudah ada
	if !statusExecute {
		// update qty order dan total order
		orderQty := order.Qty + input.Qty
		orderTotal := priceProduct * orderQty
		// set qty order dan total order
		order.Total = orderTotal
		order.Qty = orderQty
		order.UpdatedAt = time.Now()
		// update data order
		err := db.Save(&order).Error
		if err != nil {
			return product, order, errors.New("Product gagal ditambahkan di cart")
		}
	} else {
		// set data order
		order.IDProduct = idProduct
		order.IDUser = userId
		order.TglOrder = time.Now()
		order.Status = models.StatusAllowed("cart")
		// set total order dan qty
		order.Total = input.Qty * priceProduct
		order.Qty = input.Qty
		order.CreatedAt = time.Now()
		// create data order
		err = db.Create(&order).Error
		if err != nil {
			return product, order, errors.New("Product gagal disimpan di cart")
		}
	}
	return product, order, nil
}

func OrderCheckoutService(input models.OrderCheckout, userId uint) (models.Product, models.Order, error) {
	order := models.Order{}
	product := models.Product{}
	db := database.GetDB()

	idOrder := input.IDOrder

	// cek order jika sudah ada di order, maka akan di update
	err := db.Where("id_order = ? AND id_user = ?", idOrder, userId).Find(&order).Error
	if err != nil {
		return product, order, errors.New("Data order tidak ditemukkan")
	}

	// set id product dan status checkout
	idProduct := order.IDProduct
	statusCheckout := order.Status

	// buat kondisi status checkout
	if statusCheckout == "checkout" {
		// set jika order sudah pernah di checkout
		return product, order, errors.New("Data order sudah pernah di checkout! Anda tidak bisa melakukan checkout.")
	} else if statusCheckout == "cart" {
		// cari product
		err := db.Where("id_product = ?", idProduct).Find(&product).Error
		if err != nil {
			return product, order, errors.New("Product tidak ditemukkan")
		}

		// set kurang stock di product berdasarkan order qty
		stockProduct := product.Stock - order.Qty
		product.Stock = stockProduct
		// set update product time
		product.UpdatedAt = time.Now()
		// update stok product
		err = db.Save(&product).Error
		if err != nil {
			return product, order, errors.New("Gagal mengupdate stock product")
		}

		//set data order checkout
		order.Status = models.StatusAllowed("checkout")
		order.TglCheckout = time.Now()
		order.UpdatedAt = time.Now()

		// update status dan tanggal checkout
		err = db.Save(&order).Error
		if err != nil {
			return product, order, err
		}
	} else if statusCheckout == "cancel" {
		// set jika order sudah pernah di cancel
		return product, order, errors.New("Data order sudah pernah di cancel! Anda tidak bisa melakukan checkout.")
	}
	return product, order, nil
}

func OrderCancelService(input models.OrderCancel, userId uint) (models.Product, models.Order, error) {
	order := models.Order{}
	product := models.Product{}
	db := database.GetDB()

	idOrder := input.IDOrder

	// cek order jika sudah ada di order, maka akan di update
	err := db.Where("id_order = ? AND id_user = ?", idOrder, userId).Find(&order).Error
	if err != nil {
		return product, order, errors.New("Data order tidak ditemukkan")
	}

	// set id product dan status
	idProduct := order.IDProduct
	statusCheckout := order.Status

	// set data order cancel
	order.Status = models.StatusAllowed("cancel")
	order.TglCancel = time.Now()
	order.UpdatedAt = time.Now()

	// buat kondisi status
	if statusCheckout == "cancel" {
		// set jika order sudah pernah di cancel
		return product, order, errors.New("Data order sudah pernah di cancel! Anda tidak bisa melakukan cancel.")
	} else if statusCheckout == "cart" {

		// update order status dan tanggal cancel
		err := db.Save(&order).Error
		if err != nil {
			return product, order, err
		}
	} else if statusCheckout == "checkout" {
		// kita akan rollback stock yang sudah berkurang
		// cari product
		err := db.Where("id_product = ?", idProduct).Find(&product).Error
		if err != nil {
			return product, order, errors.New("Product tidak ditemukkan")
		}

		// set kembalikan stock semula berdasarkan order qty
		stockProduct := product.Stock + order.Qty
		product.Stock = stockProduct
		// set update product time
		product.UpdatedAt = time.Now()
		// update stock product
		err = db.Save(&product).Error
		if err != nil {
			return product, order, errors.New("Gagal mengembalikan stock semula")
		}

		// update order status dan tanggal cancel
		err = db.Save(&order).Error
		if err != nil {
			return product, order, err
		}
	}

	return product, order, nil
}

func OrderGetServiceByIdUser(userId uint) ([]models.Order, error) {
	var order []models.Order
	db := database.GetDB()

	err := db.Preload("Product").Where("id_user = ?", userId).Find(&order).Error
	if err != nil {
		return order, errors.New("Data order tidak ditemukkan")
	}
	return order, nil
}

func OrderGetServiceByIdOrder(userId uint, idOrder uint) (models.Order, error) {
	var order models.Order
	db := database.GetDB()

	// cari order
	err := db.Preload("Product").Where("id_user = ? AND id_order = ?", userId, idOrder).Find(&order).Error
	if err != nil {
		return order, errors.New("Data order tidak ditemukkan")
	}

	return order, nil
}
