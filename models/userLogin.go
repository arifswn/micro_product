package models

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

func FormatResponseLogin(token string) LoginUserResponse {
	result := LoginUserResponse{
		Token: token,
	}
	return result
}
