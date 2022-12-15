package routers

import (
	"micro_product/controllers"
	"micro_product/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	orderRouter := router.Group("/api/v1/orders")
	{
		orderRouter.Use(middlewares.Authentication())
		// Order add to cart
		orderRouter.POST("/cart", controllers.OrderCart)
		// Order checkout product
		orderRouter.POST("/checkout", controllers.OrderCheckout)
		// Order cancel product
		orderRouter.POST("/cancel", controllers.OrderCancel)
		// Get data order by id user
		orderRouter.GET("", controllers.OrderGetByIDUser)
		// Get data order by id order
		orderRouter.GET("/:orderId", controllers.OrderGetByIDOrder)
	}
}
