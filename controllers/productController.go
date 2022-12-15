package controllers

import (
	"micro_product/helpers"
	"micro_product/models"
	"micro_product/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary			Create Product
// @Description		Digunakan untuk melakukan create product
// @Tags 			Products
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			Product		body		models.ProductInput				true	"Create Product"
// @Success 		201			{object} 	models.ProductInputResponse
// @Failure      	400  		{object}  	models.ProductInputResponse
// @Failure      	404  		{object}  	models.ProductInputResponse
// @Failure      	422  		{object}  	models.ProductInputResponse
// @Failure      	500  		{object}  	models.ProductInputResponse
// @Router 			/products 	[post]
func CreateProduct(c *gin.Context) {
	var input models.ProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Product created failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := services.ProductInputService(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Product created failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseProductInput(newProduct)
	response := helpers.APIResponse("Product created successfully", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

// UpdateProduct godoc
// @Summary			Update Product
// @Description		Digunakan untuk melakukan update product
// @Tags 			Products
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			productId 	path 		int 							true	"ID"
// @Param 			Product		body		models.ProductUpdate			true	"Update Product"
// @Success 		200			{object} 	models.ProductUpdateResponse
// @Failure      	400  		{object}  	models.ProductUpdateResponse
// @Failure      	404  		{object}  	models.ProductUpdateResponse
// @Failure      	422  		{object}  	models.ProductUpdateResponse
// @Failure      	500  		{object}  	models.ProductUpdateResponse
// @Router 			/products/{productId} 	[put]
func UpdateProduct(c *gin.Context) {
	var input models.ProductUpdate
	productId, _ := strconv.Atoi(c.Param("productId"))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Product updated failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateProduct, err := services.ProductUpdateService(input, uint(productId))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Product updated failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseProductUpdate(updateProduct)
	response := helpers.APIResponse("Product updated successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GetProduct godoc
// @Summary 		Get All Product
// @Description 	Digunakan untuk melakukan get all product
// @Tags 			Products
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Success 		200			{object} 	models.ProductGetResponse
// @Failure      	400  		{object}  	models.ProductGetResponse
// @Failure      	404  		{object}  	models.ProductGetResponse
// @Failure      	422  		{object}  	models.ProductGetResponse
// @Failure      	500  		{object}  	models.ProductGetResponse
// @Router 			/products 	[get]
func GetProduct(c *gin.Context) {
	product, err := services.ProductGetService()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helpers.APIResponse("Get Product failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseProductGet(product)
	response := helpers.APIResponse("Get Product successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// GetProductById godoc
// @Summary 		Get Product By ID Product
// @Description 	Digunakan untuk melakukan get product by id product
// @Tags 			Products
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			productId 	path 		int 							true	"ID"
// @Success 		200			{object} 	models.ProductGetResponse
// @Failure      	400  		{object}  	models.ProductGetResponse
// @Failure      	404  		{object}  	models.ProductGetResponse
// @Failure      	422  		{object}  	models.ProductGetResponse
// @Failure      	500  		{object}  	models.ProductGetResponse
// @Router 			/products/{productId} 	[get]
func GetProductById(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	product, err := services.ProductGetServiceById(uint(productId))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helpers.APIResponse("Get Product`s failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseProduct(product)
	response := helpers.APIResponse("Get Product`s successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// DeleteProduct godoc
// @Summary			Delete Product
// @Description		Digunakan untuk melakukan delete product
// @Tags 			Products
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			productId 	path 		int 							true	"ID"
// @Success 		200			{object} 	string
// @Failure      	400  		{object}  	string
// @Failure      	404  		{object}  	string
// @Failure      	422  		{object}  	string
// @Failure      	500  		{object}  	string
// @Router 			/products/{productId} 	[delete]
func DeleteProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	IDProduct := uint(productId)

	product, err := services.ProductGetServiceById(IDProduct)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Product deleted failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if product.IDProduct != IDProduct {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Product deleted Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, err := services.ProductDeleteService(IDProduct)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Product deleted Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Product deleted successfully", http.StatusOK, "success", res)
	c.JSON(http.StatusOK, response)
}
