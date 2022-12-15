package models

import "time"

type RegisterUserInput struct {
	Age      uint   `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type RegisterUserResponse struct {
	IDUser    uint      `json:"id_user"`
	Age       uint      `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatResponseRegister(user User) RegisterUserResponse {
	result := RegisterUserResponse{
		IDUser:    user.IDUser,
		Age:       user.Age,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
	return result
}
