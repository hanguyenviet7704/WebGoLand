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
)

type RolesController interface {
	GetAllRoles(ctx *gin.Context)
	GetRoleById(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	GetPermissionsFromRole(ctx *gin.Context)
	AssignPermissionsToRole(ctx *gin.Context)
}
type rolesController struct {
	roleService service.RoleService
}

func NewRolesController(roleService service.RoleService) RolesController {
	return &rolesController{
		roleService: roleService,
	}
}

func (c *rolesController) GetAllRoles(ctx *gin.Context) {
	roles, err := c.roleService.GetAllRoles()
	if err != nil {
		global.Logger.Error("Lỗi khi lấy danh sách roles: ", zap.Error(err))
		response.ErrorResponse(ctx, 40001)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, gin.H{"roles": roles})
}

func (c *rolesController) GetRoleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	role, err := c.roleService.GetRoleByID(id)
	if err != nil {
		global.Logger.Error("Lỗi khi lấy role theo ID: ", zap.Error(err))
		response.ErrorResponse(ctx, 40002)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, gin.H{"role": role})
}

func (c *rolesController) CreateRole(ctx *gin.Context) {
	var role po.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		global.Logger.Error("Lỗi khi bind JSON tạo role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40003)
		return
	}
	if err := c.roleService.CreateRole(&role); err != nil {
		global.Logger.Error("Lỗi khi tạo role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40004)
		return
	}
	response.SuccessResponse(ctx, http.StatusCreated, gin.H{"role": role})
}

func (c *rolesController) UpdateRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	var role po.Roles
	if err := ctx.ShouldBindJSON(&role); err != nil {
		global.Logger.Error("Lỗi khi bind JSON cập nhật role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40003)
		return
	}
	role.ID = id
	if err := c.roleService.UpdateRole(&role); err != nil {
		global.Logger.Error("Lỗi khi cập nhật role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40004)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, gin.H{"role": role})
}

func (c *rolesController) DeleteRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	if err := c.roleService.DeleteRole(id); err != nil {
		global.Logger.Error("Lỗi khi xóa role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40005)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, gin.H{"message": "Role deleted", "id": id})
}

func (c *rolesController) GetPermissionsFromRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	permissions, err := c.roleService.GetPermissionsFromRole(id)
	if err != nil {
		global.Logger.Error("Lỗi khi lấy permissions theo role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40006)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, gin.H{"permissions": permissions})
}

func (c *rolesController) AssignPermissionsToRole(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, 40001)
		return
	}
	var input struct {
		PermissionIDs []uint `json:"permissionIds"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		global.Logger.Error("Lỗi khi bind JSON phân quyền role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40003)
		return
	}
	if err := c.roleService.AssignPermissionsToRole(id, input.PermissionIDs); err != nil {
		global.Logger.Error("Lỗi khi phân quyền role: ", zap.Error(err))
		response.ErrorResponse(ctx, 40007)
		return
	}
	response.SuccessResponse(ctx, http.StatusOK, gin.H{"message": "Permissions assigned", "roleId": id})
}
