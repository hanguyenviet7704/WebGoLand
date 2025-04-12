package middlewares

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type JWTMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
	RoleMiddleware(allowedRoles ...string) gin.HandlerFunc
	PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc
}

type jwtMiddleware struct {
	jwtService service.JWTService
}

func NewJWTMiddleware(jwtService service.JWTService) JWTMiddleware {
	return &jwtMiddleware{
		jwtService: jwtService,
	}
}

func (m *jwtMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			global.Logger.Warn("Authorization header rỗng")
			response.ErrorResponse(ctx, 20002)
			ctx.Abort()
			return
		}
		const BEARER_SCHEMA = "Bearer "
		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			global.Logger.Warn("Lỗi định dạng Authorization")
			response.ErrorResponse(ctx, 20002)
			ctx.Abort()
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			global.Logger.Error("Lỗi xác thực token", zap.Error(err))
			response.ErrorResponse(ctx, 30001)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
		return
	}
}

func (m *jwtMiddleware) RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Get("claims")
		if !ok {
			global.Logger.Warn("Không tìm thấy claims trong context")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}

		authClaims, ok := claims.(*service.AuthCustomClaims)
		if !ok {
			global.Logger.Error("Kiểu dữ liệu claims không hợp lệ")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}

		for _, userRole := range authClaims.Roles {
			for _, allowed := range allowedRoles {
				if allowed == userRole.Name {
					ctx.Next()
					return
				}
			}
		}

		global.Logger.Warn("Không có vai trò phù hợp để truy cập", zap.Strings("required_roles", allowedRoles))
		response.ErrorResponse(ctx, 30004)
		ctx.Abort()
		return
	}
}

func (m *jwtMiddleware) PermissionMiddleware(allowedPermissions ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Get("claims")
		if !ok {
			global.Logger.Error("Không tìm thấy claims trong context")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}

		authClaims, ok := claims.(*service.AuthCustomClaims)
		if !ok {
			global.Logger.Error("Kiểu dữ liệu claims không hợp lệ")
			response.ErrorResponse(ctx, 40001)
			ctx.Abort()
			return
		}
		for _, userPerm := range authClaims.Permissions {
			for _, allowed := range allowedPermissions {
				if allowed == userPerm.Name {
					ctx.Next()
					return
				}
			}
		}

		global.Logger.Warn("Không có quyền phù hợp để truy cập", zap.Strings("required_permissions", allowedPermissions))
		response.ErrorResponse(ctx, 30004)
		ctx.Abort()
		return
	}
}
