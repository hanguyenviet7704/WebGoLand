package service

import (
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/repository"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserService interface {
	LoginUser(email string, password string) (*po.User, string, string, error)
	RegisterUser(username string, email string, password string) (*po.User, error)

	GetAllUsers() (*[]po.User, error)
	GetUserById(id int) (*po.User, error)
	CreateOrUpdateUser(user *po.User, roleName []string) error
	DeleteUser(id int) error
	GetRoleFromUser(id int) (*[]po.Roles, error)
	AssignRolesToUser(id int, roleNames []string) error
}
type userService struct {
	userRepository  repository.UserRepository
	jwtService      JWTService
	tokenRepository repository.TokenRepository
}

func NewUserService(userRepo repository.UserRepository, jwtServices JWTService, tokenRepository repository.TokenRepository) UserService {
	return &userService{
		tokenRepository: tokenRepository,
		userRepository:  userRepo,
		jwtService:      jwtServices,
	}
}
func (service *userService) LoginUser(email string, password string) (*po.User, string, string, error) {
	user, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("USER_NOT_FOUND")
	}
	if user.Password != password {
		return nil, "", "", errors.New("WRONG_PASSWORD")
	}
	accessToken, err := service.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, "", "", errors.New("ACCESS_TOKEN_ERROR")
	}
	refreshToken, err := service.jwtService.GenerateRefreshToken(user)
	if err != nil {
		return nil, "", "", errors.New("REFRESH_TOKEN_ERROR")
	}
	err = service.tokenRepository.SaveToken(user, accessToken, refreshToken)
	if err != nil {
		return nil, "", "", errors.New("TOKEN_ERROR")
	}
	return user, accessToken, refreshToken, nil
}
func (service *userService) RegisterUser(name, email, password string) (*po.User, error) {
	existingUser, _ := service.userRepository.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("EMAIL_EXIST")
	}
	user := &po.User{
		Name:      name,
		Email:     email,
		Password:  password,
		Is_active: true,
	}

	err := service.userRepository.CreateUserWithDefaultRoles(user)
	if err != nil {
		return nil, errors.New("DATABASE_ERROR")
	}
	return user, nil
}
func (service *userService) GetAllUsers() (*[]po.User, error) {
	users, err := service.userRepository.FindAll()
	return users, err
}
func (service *userService) GetUserById(id int) (*po.User, error) {
	user, err := service.userRepository.FindUserById(id)
	return user, err
}
func (service *userService) CreateOrUpdateUser(user *po.User, roleNames []string) error {
	//Email

	// Vai trò
	var roles []po.Roles
	roles, err := service.userRepository.FindRolesByNames(roleNames)
	if err != nil {
		return errors.New("ROLE_ERROR")
	}
	if len(roles) == 0 {
		return errors.New("ROLE_NOT_FOUND")
	}

	user.Roles = roles
	var existingUser *po.User
	existingUser, err = service.userRepository.FindUserById(user.Id)
	if err != nil {
		//User không tồn tại ==> tạo mới
		user.CreatedAt = time.Now()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := service.userRepository.Create(user); err != nil {
				return errors.New("ERROR_CREATE_USER")
			}
		}
	} else {
		// User đã tồn tại => cập nhật
		user.CreatedAt = existingUser.CreatedAt
		if err := service.userRepository.Save(user); err != nil {
			return errors.New("ERROR_SAVE_USER")
		}
	}
	return nil
}
func (service *userService) DeleteUser(id int) error {
	return service.userRepository.Delete(&po.User{Id: id})
}
func (service *userService) GetRoleFromUser(id int) (*[]po.Roles, error) {
	_, err := service.userRepository.FindUserById(id)
	if err != nil {
		return nil, errors.New("Không tìm thấy user")
	}
	return service.userRepository.FindRoleFromUser(id)
}
func (service *userService) AssignRolesToUser(id int, roleNames []string) error {
	roles := []po.Roles{}
	roles, err := service.userRepository.FindRolesByNames(roleNames)
	if err != nil {
		return errors.New("Không thể tìm thấy Role trong database")
	}
	err = service.userRepository.CreateRoleFromUser(id, &roles)
	if err != nil {
		return err
	}
	return nil
}
