package routers

import (
	"encoding/json"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	// Disable Console Color
	// gin.Disable ConsoleColor()
	// r := gin.Default() //change default

	r := gin.New()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(gin.Recovery())         // to recover gin automatically
	r.Use(jsonLoggerMiddleware()) // we'll define it later
	r.Use(location.Default())

	// call to browser http://localhost:8080/swagger/index.html
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthRoutes(r)
	ProductRoutes(r)
	OrderRoutes(r)

	return r
}

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}
