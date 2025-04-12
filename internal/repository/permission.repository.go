package repository

import (
	"awesomeProject5/internal/po"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindAll() ([]po.Permissions, error)
	FindById(id string) (po.Permissions, error)
	Create(permission *po.Permissions) error
	Update(permission *po.Permissions) error
	Delete(id string) error
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) FindAll() ([]po.Permissions, error) {
	var permissions []po.Permissions
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindById(id string) (po.Permissions, error) {
	var permission po.Permissions
	err := r.db.First(&permission, id).Error
	return permission, err
}

func (r *permissionRepository) Create(permission *po.Permissions) error {
	return r.db.Create(&permission).Error
}

func (r *permissionRepository) Update(permission *po.Permissions) error {
	return r.db.Save(&permission).Error
}

func (r *permissionRepository) Delete(id string) error {
	return r.db.Delete(&po.Permissions{}, id).Error
}
