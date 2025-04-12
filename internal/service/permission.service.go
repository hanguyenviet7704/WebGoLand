package service

import (
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/repository"
)

type PermissionService interface {
	GetAllPermissions() ([]po.Permissions, error)
	GetPermissionById(id string) (po.Permissions, error)
	CreatePermission(permission *po.Permissions) error
	UpdatePermission(id string, permission *po.Permissions) error
	DeletePermission(id string) error
}

type permissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionService{repo: repo}
}

func (s *permissionService) GetAllPermissions() ([]po.Permissions, error) {
	return s.repo.FindAll()
}

func (s *permissionService) GetPermissionById(id string) (po.Permissions, error) {
	return s.repo.FindById(id)
}

func (s *permissionService) CreatePermission(permission *po.Permissions) error {
	return s.repo.Create(permission)
}

func (s *permissionService) UpdatePermission(id string, permission *po.Permissions) error {
	// có thể kiểm tra trước khi update, ví dụ như validate dữ liệu.
	return s.repo.Update(permission)
}

func (s *permissionService) DeletePermission(id string) error {
	return s.repo.Delete(id)
}
