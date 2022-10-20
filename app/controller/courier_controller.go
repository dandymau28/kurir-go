package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	GetAllCourier(ctx *gin.Context)
}
