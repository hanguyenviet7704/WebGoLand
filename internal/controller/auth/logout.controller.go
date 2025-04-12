package auth

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type LogoutController interface {
	Logout(ctx *gin.Context)
	LogoutAllDevices(ctx *gin.Context)
}
type logoutController struct {
	tokenService service.TokenService
}

func NewLogout(tokenService service.TokenService) LogoutController {
	return &logoutController{
		tokenService: tokenService,
	}
}
func (controller *logoutController) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		global.Logger.Warn("Authorization header rỗng")
		response.ErrorResponse(ctx, 20002)
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")
	err := controller.tokenService.DeleteAccessToken(token)
	if err != nil {
		global.Logger.Warn("Lỗi khi xóa token", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}

	global.Logger.Info("Đăng xuất thành công thiết bị")
	response.SuccessResponse(ctx, 20001, "Đăng xuất khỏi thiết bị thành công ")
	return
}

func (controller *logoutController) LogoutAllDevices(ctx *gin.Context) {
	var request struct {
		UserID int `json:"user_id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil || request.UserID == 0 {
		global.Logger.Warn("User Id không hợp lệ")
		response.ErrorResponse(ctx, 20003)
		return
	}
	err := controller.tokenService.DeleteAccessTokenByUserId(request.UserID)
	if err != nil {
		global.Logger.Error("Lỗi khi xóa tất cả token", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}
	global.Logger.Info("Đăng xuất khỏi tất cả thiết bị thành công", zap.Int("user_id", request.UserID))
	response.SuccessResponse(ctx, 20001, "Đăng xuất khỏi tất cả thiết bị thành công")
	return
}
