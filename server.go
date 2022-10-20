package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kurir-go/app/controller"
	"github.com/kurir-go/app/repository"
	"github.com/kurir-go/app/service"
	"github.com/kurir-go/config"
	"github.com/kurir-go/middleware"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.DatabaseConnection()
	userRepository     repository.UserRepository     = repository.NewUserRepository(db)
	shipmentRepository repository.ShipmentRepository = repository.NewShipmentRepository(db)
	authService        service.AuthService           = service.NewAuthService(userRepository)
	authController     controller.AuthController     = controller.NewAuthController(authService)
	userService        service.UserService           = service.NewUserService(userRepository)
	userController     controller.UserController     = controller.NewUserController(userService)
	shipmentService    service.ShipmentService       = service.NewShipmentService(shipmentRepository, userRepository)
	shipmentController controller.ShipmentController = controller.NewShipmentController(shipmentService)
)

func main() {
	defer config.CloseConnection(db)
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/users", middleware.AuthorizeJWT(authService))
	{
		userRoutes.GET("/courier", userController.GetAllCourier)
	}

	shipmentRoutes := r.Group("api/shipments", middleware.AuthorizeJWT(authService))
	{
		shipmentRoutes.GET("/", shipmentController.GetAllShipment)
		shipmentRoutes.POST("/add", shipmentController.AddShipment)
		shipmentRoutes.GET("/courier/:courier_name", shipmentController.GetCourierShipment)
		shipmentRoutes.PUT("/update/:shipment_id/delivery", shipmentController.UpdateDelivery)
		shipmentRoutes.GET("/:shipment_id/detail", shipmentController.GetShipmentDetail)
	}

	r.Run()

}
