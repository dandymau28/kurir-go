package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/service"
	"github.com/kurir-go/helper"
)

type ShipmentControllerImpl struct {
	shipmentService service.ShipmentService
}

func NewShipmentController(shipmentService service.ShipmentService) ShipmentController {
	return &ShipmentControllerImpl{
		shipmentService: shipmentService,
	}
}

func (c *ShipmentControllerImpl) AddShipment(ctx *gin.Context) {
	var Request dto.AddShipmentRequest

	errDTO := ctx.ShouldBind(&Request)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to save shipment", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	courier_name := ctx.GetString("username")

	errService := c.shipmentService.AddShipment(Request.ClientName, Request.ClientAddress, courier_name, Request.AWBNo)
	if errService != nil {
		response := helper.BuildErrorResponse("Failed to save shipment", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *ShipmentControllerImpl) GetAllShipment(ctx *gin.Context) {
	result := c.shipmentService.GetAllShipment()

	response := helper.BuildResponse(true, "data found", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *ShipmentControllerImpl) GetCourierShipment(ctx *gin.Context) {
	courierParam := ctx.Param("courier_name")
	courierName := ctx.GetString("username")
	userRole := ctx.GetString("role")

	if courierName != courierParam && userRole == "kurir" {
		response := helper.BuildErrorResponse("Cannot access this data", "forbidden", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if courierName == "" {
		response := helper.BuildErrorResponse("Failed to get data", "no courier found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.shipmentService.GetShipmentByCourier(courierName)

	if err != nil {
		response := helper.BuildErrorResponse("Failed to get data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "data found", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *ShipmentControllerImpl) UpdateDelivery(ctx *gin.Context) {
	paramID := ctx.Param("shipment_id")
	if paramID == "" {
		response := helper.BuildErrorResponse("Failed to update", "no shipment found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	shipmentID, _ := strconv.ParseUint(paramID, 0, 64)

	errService := c.shipmentService.UpdateDeliveryShipment(shipmentID)
	if errService != nil {
		response := helper.BuildErrorResponse("Failed to update", errService.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "success", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *ShipmentControllerImpl) GetShipmentDetail(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID == "" {
		response := helper.BuildErrorResponse("Failed to get data", "no shipment found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	shipmentID, _ := strconv.ParseUint(paramID, 0, 64)

	result := c.shipmentService.GetShipmentDetailByID(shipmentID)

	response := helper.BuildResponse(true, "data found", result)
	ctx.JSON(http.StatusOK, response)
}
