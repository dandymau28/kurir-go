package dto

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token       string    `json:"token"`
	TokenExpire time.Time `json:"token_expire"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
}
