package services

import (
	"errors"
	"micro_product/database"
	"micro_product/helpers"
	"micro_product/models"
	"time"
)

func UserRegisterService(input models.RegisterUserInput) (models.User, error) {
	user := models.User{}
	db := database.GetDB()

	user.Age = input.Age
	user.Email = input.Email

	// cek password
	passCheck := len(input.Password)
	if passCheck < 6 {
		return user, errors.New("Password must be at least 6 characters long")
	}

	// enc password
	passwordHash := helpers.HashPass(input.Password)
	user.Password = string(passwordHash)

	user.Username = input.Username
	user.CreatedAt = time.Now()

	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func UserLoginService(input models.LoginUserInput) (string, error) {
	user := models.User{}
	db := database.GetDB()

	//mencari user dengan email yang diinputkan
	err := db.Debug().Where("email = ?", input.Email).Find(&user).Error
	if err != nil {
		return "User not found", err
	}

	if user.IDUser == 0 {
		return "error: ", errors.New("No user found on that email")
	}

	//mencocokan password
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(input.Password))
	if !comparePass {
		return "Password not match", err
	}

	token, _ := helpers.GenerateToken(user.IDUser, user.Email)
	return token, err
}
