package initialize

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(app *AppContainer) *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	// ========================= Swagger =========================
	// Nếu có tích hợp Swagger, có thể thêm ở đây
	// ========================= Auth APIs =========================
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/login", app.LoginController.Login)
		authRoutes.POST("/refresh", app.LoginController.RefreshToken)
		authRoutes.POST("/register", app.RegisterController.Register)
		// Các route sau cần xác thực bằng JWT
		authRoutes.Use(app.MiddlewareAuth.AuthorizeJWT())
		authRoutes.POST("/logout", app.LogoutController.Logout)
		authRoutes.POST("/logoutalldevices", app.LogoutController.LogoutAllDevices)
	}
	// ========================= Admin APIs =========================
	adminRoutes := server.Group("/api/admin")
	adminRoutes.Use(
		app.MiddlewareAuth.AuthorizeJWT(),
		app.MiddlewareAuth.RoleMiddleware("ROLE_ADMIN"),
	)
	{
		// Users
		adminRoutes.GET("/users", app.UserController.GetAllUsers)
		adminRoutes.GET("/user", app.UserController.GetUserById)
		adminRoutes.POST("/users", app.UserController.CreateOrUpdateUser)
		adminRoutes.DELETE("/users", app.UserController.DeleteUser)
		adminRoutes.GET("/users/:id/roles", app.UserController.GetRolesFromUser)
		adminRoutes.PUT("/users/:id/roles", app.UserController.AssignRolesToUser)

		// Roles
		adminRoutes.GET("/roles", app.RoleController.GetAllRoles)
		adminRoutes.GET("/roles/:id", app.RoleController.GetRoleById)
		adminRoutes.POST("/roles", app.RoleController.CreateRole)
		adminRoutes.PUT("/roles/:id", app.RoleController.UpdateRole)
		adminRoutes.DELETE("/roles/:id", app.RoleController.DeleteRole)
		adminRoutes.GET("/roles/:id/permissions", app.RoleController.GetPermissionsFromRole)
		adminRoutes.PUT("/roles/:id/permissions", app.RoleController.AssignPermissionsToRole)

		// Permissions
		adminRoutes.GET("/permissions", app.PermissionController.GetAllPermissions)
		adminRoutes.GET("/permissions/:id", app.PermissionController.GetPermissionById)
		adminRoutes.POST("/permissions", app.PermissionController.CreatePermission)
		adminRoutes.PUT("/permissions/:id", app.PermissionController.UpdatePermission)
		adminRoutes.DELETE("/permissions/:id", app.PermissionController.DeletePermission)
	}

	// ========================= Protected APIs =========================
	protected := server.Group("/api")
	protected.Use(app.MiddlewareAuth.AuthorizeJWT())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Thông tin người dùng đã đăng nhập"})
		})

		protected.GET("/products",
			app.MiddlewareAuth.PermissionMiddleware("view_products"),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Bạn có quyền xem danh sách sản phẩm"})
			},
		)

		protected.GET("/dashboard",
			app.MiddlewareAuth.RoleMiddleware("ROLE_ADMIN", "ROLE_MANAGER"),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Chào bạn, đây là dashboard dành cho admin hoặc manager"})
			},
		)
	}

	return server
}
