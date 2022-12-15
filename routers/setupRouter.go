package routers

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	// Disable Console Color
	// gin.Disable ConsoleColor()
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(location.Default())

	// call to browser http://localhost:8080/swagger/index.html
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthRoutes(r)
	ProductRoutes(r)
	OrderRoutes(r)

	return r
}
