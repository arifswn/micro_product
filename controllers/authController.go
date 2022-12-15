package controllers

import (
	"micro_product/helpers"
	"micro_product/models"
	"micro_product/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister godoc
// @Summary			Registration User
// @Description		Digunakan untuk melakukan register user
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Param 			User	body			models.RegisterUserInput	true	"Register User"
// @Success 		201		{object} 		models.RegisterUserResponse
// @Failure      	400  	{object}  		models.RegisterUserResponse
// @Failure      	404  	{object}  		models.RegisterUserResponse
// @Failure      	422  	{object}  		models.RegisterUserResponse
// @Failure      	500  	{object}  		models.RegisterUserResponse
// @Router 			/users/register 	[post]
func UserRegister(c *gin.Context) {
	var input models.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := services.UserRegisterService(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseRegister(newUser)
	response := helpers.APIResponse("Account registered successfully", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

// UserLogin godoc
// @Summary 		Login User
// @Description 	Digunakan untuk melakukan login user
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Param 			User	body			models.LoginUserInput		true	"Login User"
// @Success			200 	{object} 		models.LoginUserResponse
// @Failure      	400  	{object}  		models.RegisterUserResponse
// @Failure      	404  	{object}  		models.RegisterUserResponse
// @Failure      	422  	{object}  		models.RegisterUserResponse
// @Failure      	500  	{object}  		models.RegisterUserResponse
// @Router 			/users/login [post]
func UserLogin(c *gin.Context) {
	var input models.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := services.UserLoginService(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := models.FormatResponseLogin(loginUser)
	response := helpers.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
