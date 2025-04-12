package service

import (
	"awesomeProject5/internal/repository"
)

type TokenService interface {
	FindRefreshToken(refreshToken string) error
	UpdateAccessToken(accessToken string, refreshToken string) error
	DeleteAccessToken(accessToken string) error
	DeleteAccessTokenByUserId(userId int) error
}
type tokenService struct {
	tokenRepository repository.TokenRepository
}

func NewTokenService(tokenRepository repository.TokenRepository) TokenService {
	return &tokenService{
		tokenRepository: tokenRepository,
	}
}
func (service *tokenService) FindRefreshToken(refreshToken string) error {
	err := service.tokenRepository.FindRefreshToken(refreshToken)
	return err
}
func (service *tokenService) UpdateAccessToken(accessToken string, refreshToken string) error {
	err := service.tokenRepository.UpdateToken(accessToken, refreshToken)
	return err
}
func (service *tokenService) DeleteAccessToken(accessToken string) error {
	err := service.tokenRepository.DeleteToken(accessToken)
	return err
}
func (service *tokenService) DeleteAccessTokenByUserId(userId int) error {
	err := service.tokenRepository.DeleteTokenByUserId(userId)
	return err
}
