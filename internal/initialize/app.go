package initialize

import (
	"awesomeProject5/internal/controller"
	"awesomeProject5/internal/controller/auth"
	"awesomeProject5/internal/middlewares"
	"awesomeProject5/internal/repository"
	"awesomeProject5/internal/service"
	"gorm.io/gorm"
)

type AppContainer struct {
	UserService          service.UserService
	JWTService           service.JWTService
	LoginController      auth.LoginController
	LogoutController     auth.LogoutController
	RegisterController   auth.RegisterController
	UserController       controller.UserController
	RoleController       controller.RolesController
	PermissionController controller.PermissionController
	MiddlewareAuth       middlewares.JWTMiddleware
	TokenRepository      repository.TokenRepository
	TokenService         service.TokenService
}

func InitApp(db *gorm.DB) *AppContainer {
	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	jwtService := service.NewJWTService(userRepo)
	userService := service.NewUserService(userRepo, jwtService, tokenRepo)
	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	middlewareAuth := middlewares.NewJWTMiddleware(jwtService)
	tokenService := service.NewTokenService(tokenRepo)
	return &AppContainer{
		UserService:          userService,
		JWTService:           jwtService,
		LoginController:      auth.NewLoginController(tokenService, userService, jwtService),
		LogoutController:     auth.NewLogout(tokenService),
		RegisterController:   auth.NewRegisterController(userService),
		UserController:       controller.NewUserController(userService),
		RoleController:       controller.NewRolesController(roleService),
		PermissionController: controller.NewPermissionController(db),
		MiddlewareAuth:       middlewareAuth,
	}
}
