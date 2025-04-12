package repository

import (
	"awesomeProject5/internal/po"
	"errors"
	"gorm.io/gorm"
	"time"
)

type TokenRepository interface {
	SaveToken(user *po.User, accessToken string, refreshToken string) error
	FindRefreshToken(refreshToken string) error
	UpdateToken(accessToken string, refreshToken string) error
	DeleteToken(accessToken string) error
	DeleteTokenByUserId(id int) error
}
type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (r *tokenRepository) SaveToken(user *po.User, accessToken string, refreshToken string) error {
	result := r.db.Create(&po.Tokens{
		UserID:                   user.Id,
		AccessToken:              accessToken,
		RefreshToken:             refreshToken,
		Access_token_expires_at:  time.Now().Add(time.Minute * 60),
		Refresh_token_expires_at: time.Now().Add(time.Hour * 24),
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	})
	if result.Error != nil {
		return errors.New("Không thể lưu Token xuống database")
	}
	return nil
}
func (r *tokenRepository) FindRefreshToken(refreshToken string) error {
	var token po.Tokens
	err := r.db.Where("refresh_token = ?", refreshToken).First(&token).Error
	return err
}

func (r *tokenRepository) UpdateToken(accessToken string, refreshToken string) error {
	err := r.db.Model(&po.Tokens{}).
		Where("refresh_token = ?", refreshToken).
		Update("access_token", accessToken).Error
	return err
}
func (r *tokenRepository) DeleteToken(accessToken string) error {
	result := r.db.Where("access_token = ?", accessToken).Delete(&po.Tokens{})
	if result.Error != nil {
		return errors.New("DELETE_ERROR")
	}
	if result.RowsAffected == 0 {
		return errors.New("TOKEN_NOT_FOUND")
	}
	return nil
}

func (r *tokenRepository) DeleteTokenByUserId(id int) error {
	result := r.db.Where("user_id = ?", id).Delete(&po.Tokens{})
	if result.Error != nil {
		return errors.New("DELETE_ERROR")
	}
	if result.RowsAffected == 0 {
		return errors.New("USER_ID_NOTFOUND")
	}
	return nil
}
