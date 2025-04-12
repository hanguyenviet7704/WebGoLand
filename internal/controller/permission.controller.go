package controller

import (
	"awesomeProject5/global"
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/service"
	"awesomeProject5/response"
	"github.com/gin-gonic/gin"
)

type PermissionController interface {
	GetAllPermissions(ctx *gin.Context)
	GetPermissionById(ctx *gin.Context)
	CreatePermission(ctx *gin.Context)
	UpdatePermission(ctx *gin.Context)
	DeletePermission(ctx *gin.Context)
}

type permissionController struct {
	service service.PermissionService
}

func NewPermissionController(service service.PermissionService) PermissionController {
	return &permissionController{service: service}
}

func (c *permissionController) GetAllPermissions(ctx *gin.Context) {
	permissions, err := c.service.GetAllPermissions()
	if err != nil {
		switch err.Error() {
		case "DATABASE_ERROR":
			global.Logger.Error("Lỗi truy vấn cơ sở dữ liệu")
			response.ErrorResponse(ctx, response.ErrDatabase)
		default:
			global.Logger.Error("Lỗi hệ thống khi lấy danh sách quyền")
			response.ErrorResponse(ctx, response.ErrInternalServer)
		}
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, permissions)
}

func (c *permissionController) GetPermissionById(ctx *gin.Context) {
	id := ctx.Param("id")
	permission, err := c.service.GetPermissionById(id)
	if err != nil {
		switch err.Error() {
		case "PERMISSION_NOT_FOUND":
			global.Logger.Warn("Không tìm thấy quyền với ID: " + id)
			response.ErrorResponse(ctx, response.ErrUserNotFound)
		default:
			global.Logger.Error("Lỗi hệ thống khi tìm quyền với ID: " + id)
			response.ErrorResponse(ctx, response.ErrInternalServer)
		}
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, permission)
}

func (c *permissionController) CreatePermission(ctx *gin.Context) {
	var permission po.Permissions
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		global.Logger.Warn("Yêu cầu không hợp lệ khi tạo quyền")
		response.ErrorResponse(ctx, response.ErrCodeBadRequest)
		return
	}
	if err := c.service.CreatePermission(&permission); err != nil {
		switch err.Error() {
		case "DUPLICATE_PERMISSION":
			global.Logger.Warn("Quyền đã tồn tại")
			response.ErrorResponse(ctx, response.ErrUserAlreadyExists)
		default:
			global.Logger.Error("Lỗi khi tạo quyền mới")
			response.ErrorResponse(ctx, response.ErrInternalServer)
		}
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, permission)
}

func (c *permissionController) UpdatePermission(ctx *gin.Context) {
	id := ctx.Param("id")
	var permission po.Permissions
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		global.Logger.Warn("Yêu cầu không hợp lệ khi cập nhật quyền")
		response.ErrorResponse(ctx, response.ErrCodeBadRequest)
		return
	}
	if err := c.service.UpdatePermission(id, &permission); err != nil {
		switch err.Error() {
		case "PERMISSION_NOT_FOUND":
			global.Logger.Warn("Không tìm thấy quyền với ID: " + id)
			response.ErrorResponse(ctx, response.ErrUserNotFound)
		default:
			global.Logger.Error("Lỗi khi cập nhật quyền")
			response.ErrorResponse(ctx, response.ErrInternalServer)
		}
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, permission)
}

func (c *permissionController) DeletePermission(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.DeletePermission(id)
	if err != nil {
		switch err.Error() {
		case "PERMISSION_NOT_FOUND":
			global.Logger.Warn("Không tìm thấy quyền với ID: " + id)
			response.ErrorResponse(ctx, response.ErrUserNotFound)
		default:
			global.Logger.Error("Lỗi khi xóa quyền")
			response.ErrorResponse(ctx, response.ErrInternalServer)
		}
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
