package repository

import (
	"awesomeProject5/internal/po"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]po.Roles, error)
	FindByID(id int) (po.Roles, error)
	Create(role *po.Roles) error
	Update(role *po.Roles) error
	Delete(id int) error
	GetPermissions(id int) ([]po.Permissions, error)
	AssignPermissions(role *po.Roles, permissions []po.Permissions) error
	FindPermissionsByIDs(ids []uint) ([]po.Permissions, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindAll() ([]po.Roles, error) {
	var roles []po.Roles
	err := r.db.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *roleRepository) FindByID(id int) (po.Roles, error) {
	var role po.Roles
	err := r.db.Preload("Permissions").First(&role, id).Error
	return role, err
}

func (r *roleRepository) Create(role *po.Roles) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) Update(role *po.Roles) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id int) error {
	return r.db.Delete(&po.Roles{}, id).Error
}

func (r *roleRepository) GetPermissions(id int) ([]po.Permissions, error) {
	var role po.Roles
	err := r.db.Preload("Permissions").First(&role, id).Error
	return role.Permissions, err
}

func (r *roleRepository) AssignPermissions(role *po.Roles, permissions []po.Permissions) error {
	return r.db.Model(role).Association("Permissions").Replace(permissions)
}

func (r *roleRepository) FindPermissionsByIDs(ids []uint) ([]po.Permissions, error) {
	var permissions []po.Permissions
	err := r.db.Where("id IN ?", ids).Find(&permissions).Error
	return permissions, err
}
