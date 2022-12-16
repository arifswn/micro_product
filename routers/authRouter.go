package routers

import (
	"micro_product/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	authRouter := router.Group("/api/v1/users")
	{
		// Register User
		authRouter.POST("/register", controllers.UserRegister)
		// Login User
		authRouter.POST("/login", controllers.UserLogin)
	}
}
