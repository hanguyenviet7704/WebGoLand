package service

import (
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/repository"
)

type RoleService interface {
	GetAllRoles() ([]po.Roles, error)
	GetRoleByID(id int) (po.Roles, error)
	CreateRole(role *po.Roles) error
	UpdateRole(role *po.Roles) error
	DeleteRole(id int) error
	GetPermissionsFromRole(id int) ([]po.Permissions, error)
	AssignPermissionsToRole(id int, permissionIDs []uint) error
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) GetAllRoles() ([]po.Roles, error) {
	return s.repo.FindAll()
}

func (s *roleService) GetRoleByID(id int) (po.Roles, error) {
	return s.repo.FindByID(id)
}

func (s *roleService) CreateRole(role *po.Roles) error {
	return s.repo.Create(role)
}

func (s *roleService) UpdateRole(role *po.Roles) error {
	_, err := s.repo.FindByID(role.ID)
	if err != nil {
		return err
	}
	return s.repo.Update(role)
}

func (s *roleService) DeleteRole(id int) error {
	return s.repo.Delete(id)
}

func (s *roleService) GetPermissionsFromRole(id int) ([]po.Permissions, error) {
	return s.repo.GetPermissions(id)
}

func (s *roleService) AssignPermissionsToRole(id int, permissionIDs []uint) error {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	permissions, err := s.repo.FindPermissionsByIDs(permissionIDs)
	if err != nil {
		return err
	}
	return s.repo.AssignPermissions(&role, permissions)
}
