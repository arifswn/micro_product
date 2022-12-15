package main

import (
	"log"
	"micro_product/database"
	"micro_product/routers"
	"os"

	_ "micro_product/docs" //docs is generated by swag cli, you have to import it.

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title 			Microservices Product
// @version 		1.0
// @description		This is a sample service for managing product
// @termsOfService 	http://swagger.io/terms/
// @contact.name 	API Support
// @contact.url    	http://www.swagger.io/support
// @contact.email  	support@swagger.io
// @license.name 	Apache 2.0
// @license.url 	http://www.apache.org/licenses/LICENSE-2.0.html
// @host 			localhost:8080
// @BasePath 		/api/v1
// @securityDefinitions.apikey 	BearerAuth
// @in 							header
// @name 						Authorization
func main() {
	var appConfig = AppConfig{}

	appConfig.AppName = getEnv("APP_NAME", "MicroProduct")
	appConfig.AppEnv = getEnv("APP_ENV", "production")
	appConfig.AppPort = getEnv("APP_PORT", "8080")

	database.StartDB()
	r := routers.SetupRoutes()
	r.Run(":" + appConfig.AppPort)

}
