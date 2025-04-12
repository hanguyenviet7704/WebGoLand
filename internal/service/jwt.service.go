package service

import (
	"awesomeProject5/internal/po"
	"awesomeProject5/internal/repository"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AuthCustomClaims struct {
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	UserID      int              `json:"userId"`
	Roles       []po.Roles       `json:"role"`
	Permissions []po.Permissions `json:"permissions"`
	jwt.StandardClaims
}
type JWTService interface {
	GenerateAccessToken(user *po.User) (string, error)
	GenerateRefreshToken(user *po.User) (string, error)
	ValidateToken(token string) (*AuthCustomClaims, error)
}
type jwtServices struct {
	secretKey      string
	issuer         string
	userRepository repository.UserRepository
}

func GetSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
func NewJWTService(userRepository repository.UserRepository) JWTService {
	return &jwtServices{
		userRepository: userRepository,
		secretKey:      GetSecretKey(),
		issuer:         "VietHa",
	}
}
func (service *jwtServices) GenerateAccessToken(user *po.User) (string, error) {
	var user2 *po.User
	user2, err := service.userRepository.FindUserWithRolesAndPermissionsByID(user.Id)
	if err != nil {
		return "", err
	}
	claims := &AuthCustomClaims{
		Name:        user.Name,
		Email:       user.Email,
		UserID:      user.Id,
		Roles:       user2.Roles,
		Permissions: user2.Permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
func (service *jwtServices) GenerateRefreshToken(user *po.User) (string, error) {
	claims := &AuthCustomClaims{
		UserID: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 7 * 24).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   user.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
func (service *jwtServices) ValidateToken(encodedToken string) (*AuthCustomClaims, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Không thể xác thực token: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("Token đã hết hạn")
		}
		return claims, nil
	}
	return nil, errors.New("Không thể xác thực token ")
}
