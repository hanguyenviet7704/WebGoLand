package controller

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	CreateOrUpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetRolesFromUser(ctx *gin.Context)
	AssignRolesToUser(c *gin.Context)
}
type userController struct {
	userservice service.UserService
}

func NewUserController(userservice service.UserService) UserController {
	return &userController{
		userservice: userservice,
	}
}

// GetAllUsers godoc
// @Summary Lấy danh sách người dùng
// @Description Trả về danh sách tất cả người dùng trong hệ thống
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]entity.User
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
func (controller *userController) GetAllUsers(ctx *gin.Context) {
	users, err := controller.userservice.GetAllUsers()
	if err != nil {
		global.Logger.Error("Lỗi hệ thống khi truy vấn database")
		response.ErrorResponse(ctx, 40001)
		return
	}
	if users == nil {
		global.Logger.Error("Không có bất kì User nào")
		response.ErrorResponse(ctx, 50001)
		return
	}
	global.Logger.Info("Lấy tất cả User thành công")
	response.SuccessResponse(ctx, http.StatusOK, gin.H{
		"users": users,
	})

}

// GetUserById godoc
// @Summary Lấy thông tin người dùng theo ID
// @Description Trả về thông tin chi tiết của người dùng dựa trên ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id query int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user [get]
func (controller *userController) GetUserById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		global.Logger.Warn("Thiếu ID khi truy vấn người dùng")
		response.ErrorResponse(c, 20002) // Ví dụ lỗi thiếu tham số
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("Lỗi chuyển ID sang số nguyên: ", zap.Error(err))
		response.ErrorResponse(c, 40001)
		return
	}
	user, err := controller.userservice.GetUserById(idNum)
	if err != nil {
		global.Logger.Error("Lỗi khi truy vấn user theo ID: ", zap.Error(err))
		response.ErrorResponse(c, 40002)
		return
	}
	if user == nil {
		global.Logger.Warn("Không tìm thấy user với ID: ", zap.Int("id", idNum))
		response.ErrorResponse(c, 50001)
		return
	}
	global.Logger.Info("Lấy thông tin người dùng thành công")
	response.SuccessResponse(c, 20001, gin.H{"user": user})
	return
}
func (controller *userController) CreateOrUpdateUser(ctx *gin.Context) {
	var input struct {
		ID       string   `json:"id" `
		Name     string   `json:"name"`
		Email    string   `json:"email"`
		Password string   `json:"password"`
		IsActive bool     `json:"isActive"`
		Role     []string `json:"role"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		global.Logger.Warn("Không lấy được dữ liệu ")
		response.ErrorResponse(ctx, 20002)
		return
	}
	idNum, err := strconv.Atoi(input.ID)
	if err != nil {
		global.Logger.Error("Lỗi khi chuyển chuỗi")
		response.ErrorResponse(ctx, 40001)
		return
	}
	user := &po.User{
		Id:        idNum,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Is_active: input.IsActive,
		UpdatedAt: time.Now(),
	}
	err = controller.userservice.CreateOrUpdateUser(user, input.Role)
	if err != nil {
		switch err.Error() {
		case "ROLE_ERROR", "ROLE_NOT_FOUND":
			global.Logger.Error("Không tìm thấy bất cứ vai trò nào hoặc lỗi khi get vai trò ")
			response.ErrorResponse(ctx, 40001)
			return
		default:
			global.Logger.Error("Lỗi khi tạo hoặc cập nhật user xuống database")
			response.ErrorResponse(ctx, 40002)
			return
		}
	}
	global.Logger.Info("Tạo Hoặc Cập Nhật người dùng thành công " + user.Name)
	response.SuccessResponse(ctx, 20001, gin.H{"user": user})
	return
}

func (controller *userController) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		global.Logger.Warn("Thiếu ID khi xóa người dùng")
		response.ErrorResponse(c, 20002)
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("Lỗi chuyển ID sang số nguyên: ", zap.Error(err))
		response.ErrorResponse(c, 40001)
		return
	}
	err = controller.userservice.DeleteUser(idNum)
	if err != nil {
		global.Logger.Error("Lỗi khi xóa user: ", zap.Error(err))
		response.ErrorResponse(c, 40002)
		return
	}
	global.Logger.Info("Xóa người dùng thành công")
	response.SuccessResponse(c, 20001, gin.H{"message": "Delete User Successful", "userId": idNum})
	return
}
func (controller *userController) GetRolesFromUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		global.Logger.Warn("Thiếu ID khi lấy vai trò người dùng")
		response.ErrorResponse(c, 20001)
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("Lỗi chuyển ID sang số nguyên: ", zap.Error(err))
		response.ErrorResponse(c, 40001)
		return
	}
	roles, err := controller.userservice.GetRoleFromUser(idNum)
	if err != nil {
		global.Logger.Error("Lỗi khi lấy vai trò từ người dùng: ", zap.Error(err))
		response.ErrorResponse(c, 40002)
		return
	}
	if roles == nil {
		global.Logger.Warn("Người dùng không có vai trò nào")
		response.ErrorResponse(c, 40002)
		return
	}
	global.Logger.Info("Lấy danh sách vai trò thành công")
	response.SuccessResponse(c, 20001, gin.H{
		"userId": idNum,
		"roles":  roles,
	})
	return
}
func (controller *userController) AssignRolesToUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		global.Logger.Warn("Thiếu ID khi gán vai trò cho người dùng")
		response.ErrorResponse(ctx, 20001)
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("Lỗi chuyển ID sang số nguyên: ", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}

	var input struct {
		RoleNames []string `json:"roleNames"`
	}
	if err := ctx.BindJSON(&input); err != nil {
		global.Logger.Warn("Lỗi khi bind JSON vai trò: ", zap.Error(err))
		response.ErrorResponse(ctx, 20002)
		return
	}
	err = controller.userservice.AssignRolesToUser(idNum, input.RoleNames)
	if err != nil {
		global.Logger.Error("Lỗi khi gán vai trò: ", zap.Error(err))
		response.ErrorResponse(ctx, 40002)
		return
	}

	global.Logger.Info("Gán vai trò thành công cho user")
	response.SuccessResponse(ctx, 20001, gin.H{
		"message": "Assign Role Successful",
		"userId":  idNum,
	})
}
