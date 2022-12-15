package routers

import (
	"micro_product/controllers"
	"micro_product/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	productRouter := router.Group("/api/v1/products")
	{
		productRouter.Use(middlewares.Authentication())
		// Product Get
		productRouter.GET("", controllers.GetProduct)
		// Product Get By Id
		productRouter.GET("/:productId", controllers.GetProductById)
		// Product Create
		productRouter.POST("", controllers.CreateProduct)
		// Product Update
		productRouter.PUT("/:productId", controllers.UpdateProduct)
		// Product Delete
		productRouter.DELETE("/:productId", controllers.DeleteProduct)
	}

}
