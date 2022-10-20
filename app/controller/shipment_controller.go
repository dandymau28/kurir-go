package controller

import "github.com/gin-gonic/gin"

type ShipmentController interface {
	AddShipment(ctx *gin.Context)
	GetAllShipment(ctx *gin.Context)
	GetCourierShipment(ctx *gin.Context)
	UpdateDelivery(ctx *gin.Context)
	GetShipmentDetail(ctx *gin.Context)
}
