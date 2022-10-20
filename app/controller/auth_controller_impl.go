package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/service"
	"github.com/kurir-go/helper"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (c *AuthControllerImpl) Login(ctx *gin.Context) {
	var Request dto.LoginRequest

	errDTO := ctx.ShouldBind(&Request)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Invalid input!", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Response, err := c.authService.Login(Request.Username, Request.Password)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to login", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "success", Response)
	ctx.JSON(http.StatusOK, response)
}

func (c *AuthControllerImpl) Register(ctx *gin.Context) {
	var Request dto.RegisterRequest

	errDTO := ctx.ShouldBind(&Request)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Invalid input", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err := c.authService.Register(Request.Username, Request.Name, Request.Password, Request.Role)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to register", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Response := dto.RegisterResponse{
		Username: Request.Username,
	}

	response := helper.BuildResponse(true, "register success", Response)
	ctx.JSON(http.StatusOK, response)
}
