package controllers

import (
	"micro_product/helpers"
	"micro_product/models"
	"micro_product/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// OrderCart godoc
// @Summary			Order Cart
// @Description		Digunakan untuk melakukan order cart
// @Tags 			Orders
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			Order			body		models.OrderCart				true	"Order Cart"
// @Success 		201				{object} 	models.OrderCartResponse
// @Failure      	400  			{object}  	models.OrderCartResponse
// @Failure      	404  			{object}  	models.OrderCartResponse
// @Failure      	422  			{object}  	models.OrderCartResponse
// @Failure      	500  			{object}  	models.OrderCartResponse
// @Router 			/orders/cart 	[post]
func OrderCart(c *gin.Context) {
	var input models.OrderCart
	currentUser := c.MustGet("userData").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Order add to cart validation failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, newOrderCart, err := services.OrderCartService(input, IdUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Order add to cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseOrderCart(newProduct, newOrderCart)
	response := helpers.APIResponse("Order add to cart successfully", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

// OrderCheckout godoc
// @Summary			Order Checkout
// @Description		Digunakan untuk melakukan order checkout
// @Tags 			Orders
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			Order			body		models.OrderCheckout			true	"Order Checkout"
// @Success 		200				{object} 	models.OrderCheckoutResponse
// @Failure      	400  			{object}  	models.OrderCheckoutResponse
// @Failure      	404  			{object}  	models.OrderCheckoutResponse
// @Failure      	422  			{object}  	models.OrderCheckoutResponse
// @Failure      	500  			{object}  	models.OrderCheckoutResponse
// @Router 			/orders/checkout 	[post]
func OrderCheckout(c *gin.Context) {
	var input models.OrderCheckout
	currentUser := c.MustGet("userData").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Order checkout failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, newOrderCheckout, err := services.OrderCheckoutService(input, IdUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Order checkout failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseOrderCheckout(newProduct, newOrderCheckout)
	response := helpers.APIResponse("Order checkout successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// OrderCancel godoc
// @Summary			Order Cancel
// @Description		Digunakan untuk melakukan order cancel
// @Tags 			Orders
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			Order			body		models.OrderCancel				true	"Order Checkout"
// @Success 		200				{object} 	models.OrderCancelResponse
// @Failure      	400  			{object}  	models.OrderCancelResponse
// @Failure      	404  			{object}  	models.OrderCancelResponse
// @Failure      	422  			{object}  	models.OrderCancelResponse
// @Failure      	500  			{object}  	models.OrderCancelResponse
// @Router 			/orders/cancel 	[post]
func OrderCancel(c *gin.Context) {
	var input models.OrderCancel
	currentUser := c.MustGet("userData").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Order cancel failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, newOrderCancel, err := services.OrderCancelService(input, IdUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Order cancel failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseOrderCancel(newProduct, newOrderCancel)
	response := helpers.APIResponse("Order cancel successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// OrderGetByIdOrder godoc
// @Summary 		Order Get By Id Order
// @Description 	Digunakan untuk melakukan order get by id order
// @Tags 			Orders
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Param 			orderId 			path 		int 						true	"ID"
// @Success 		200					{object} 	models.OrderGetResponse
// @Failure      	400  				{object}  	models.OrderGetResponse
// @Failure      	404  				{object}  	models.OrderGetResponse
// @Failure      	422  				{object}  	models.OrderGetResponse
// @Failure      	500  				{object}  	models.OrderGetResponse
// @Router 			/orders/{orderId} 	[get]
func OrderGetByIDOrder(c *gin.Context) {
	orderId, _ := strconv.Atoi(c.Param("orderId"))
	currentUser := c.MustGet("userData").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	order, err := services.OrderGetServiceByIdOrder(IdUser, uint(orderId))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Get data order by id order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formmater := models.FormatResponseOrder(order)
	response := helpers.APIResponse("Get data order by id order successfully", http.StatusOK, "success", formmater)
	c.JSON(http.StatusOK, response)
}

// OrderGetByIDUser godoc
// @Summary 		Order Get By Id User
// @Description 	Digunakan untuk melakukan order get by id yser
// @Tags 			Orders
// @Accept  		json
// @Produce  		json
// @Security 		BearerAuth
// @Success 		200			{object} 	models.OrderGetResponse
// @Failure      	400  		{object}  	models.OrderGetResponse
// @Failure      	404  		{object}  	models.OrderGetResponse
// @Failure      	422  		{object}  	models.OrderGetResponse
// @Failure      	500  		{object}  	models.OrderGetResponse
// @Router 			/orders 	[get]
func OrderGetByIDUser(c *gin.Context) {
	currentUser := c.MustGet("userData").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	order, err := services.OrderGetServiceByIdUser(IdUser)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Get data order by id user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formmater := models.FormatResponseOrderAll(order)
	response := helpers.APIResponse("Get data order by id user successfully", http.StatusOK, "success", formmater)
	c.JSON(http.StatusOK, response)
}
