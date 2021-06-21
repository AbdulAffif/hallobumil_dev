package api

import (
	"github.com/AbdulAffif/hallobumil_dev/api/config"
	"github.com/AbdulAffif/hallobumil_dev/api/controller"
	"github.com/AbdulAffif/hallobumil_dev/api/middleware"
	"github.com/AbdulAffif/hallobumil_dev/api/repository"
	"github.com/AbdulAffif/hallobumil_dev/api/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	dbmssgl        *gorm.DB                  = config.SetupCon("MSSQL")
	userRepository repository.UserRepository = repository.NewUserRepository(dbmssgl)

	jwtService  service.JWTService  = service.NewJWTService()
	authService service.AuthService = service.NewAuthService(userRepository)
	userService service.UserService = service.NewUserService(userRepository)

	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func Run() {
	defer config.CloseConDB(dbmssgl)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/loginEmail", authController.LoginEmail)
		authRoutes.POST("/loginPhone", authController.LoginPhone)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/ping", authController.Ping)
	}
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		//userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run("localhost:8080")

}
