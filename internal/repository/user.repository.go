package repository

import (
	"awesomeProject5/internal/po"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type UserRepository interface {
	FindRolesByNames(names []string) ([]po.Roles, error)
	FindByEmail(email string) (*po.User, error)
	FindAll() (*[]po.User, error)
	FindUserById(id int) (*po.User, error)
	Save(user *po.User) error
	Create(user *po.User) error
	Delete(user *po.User) error
	FindRoleFromUser(id int) (*[]po.Roles, error)
	CreateRoleFromUser(id int, roles *[]po.Roles) error
	FindUserWithRolesAndPermissionsByID(id int) (*po.User, error)
	CreateUserWithDefaultRoles(user *po.User) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) FindRolesByNames(names []string) ([]po.Roles, error) {
	var roles []po.Roles
	if err := r.db.Where("name IN ?", names).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
func (r *userRepository) FindByEmail(email string) (*po.User, error) {
	var user po.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("User không tồn tại")
	}
	return &user, err
}
func (r *userRepository) FindAll() (*[]po.User, error) {
	var users []po.User
	err := r.db.Preload("Roles").Find(&users).Error
	return &users, err
}
func (r *userRepository) FindUserById(id int) (*po.User, error) {
	var user po.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}
func (r *userRepository) Save(user *po.User) error {
	err := r.db.Save(user).Error
	return err
}
func (r *userRepository) Create(user *po.User) error {
	err := r.db.Create(user).Error
	return err
}
func (r *userRepository) Delete(user *po.User) error {
	if err := r.db.Model(user).Association("Roles").Clear(); err != nil {
		return err
	}
	if err := r.db.Model(user).Association("Permissions").Clear(); err != nil {
		return err
	}
	if err := r.db.Where("user_id = ?", user.Id).Delete(&po.Tokens{}).Error; err != nil {
		return err
	}
	err := r.db.Delete(user).Error
	return err
}
func (r *userRepository) FindRoleFromUser(id int) (*[]po.Roles, error) {
	var user po.User
	err := r.db.Preload("Roles").Find(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user.Roles, nil
}

func (r *userRepository) CreateRoleFromUser(id int, roles *[]po.Roles) error {
	var user po.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return errors.New("Can't find user with id " + strconv.Itoa(id))
	}
	err = r.db.Model(&user).Association("Roles").Append(roles)
	if err != nil {
		return errors.New("Can't create role for user with id " + strconv.Itoa(id))
	}
	return nil
}
func (r *userRepository) FindUserWithRolesAndPermissionsByID(id int) (*po.User, error) {
	var user po.User
	if err := r.db.Preload("Roles").Preload("Permissions").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) CreateUserWithDefaultRoles(user *po.User) error {
	var roles []po.Roles
	if err := r.db.Where("name = ?", "ROLE_USER").First(&roles).Error; err != nil {
		return err
	}
	user.Roles = roles
	err := r.db.Create(user).Error
	return err
}
