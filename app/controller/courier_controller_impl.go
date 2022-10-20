package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurir-go/app/service"
	"github.com/kurir-go/helper"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (c *UserControllerImpl) GetAllCourier(ctx *gin.Context) {
	couriers, err := c.userService.GetAll()

	if err != nil {
		response := helper.BuildErrorResponse("Fail to get couriers", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "data found", couriers)
	ctx.JSON(http.StatusOK, response)
}
