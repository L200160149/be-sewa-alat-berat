package repository

import (
	"context"

	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"gorm.io/gorm"
)
type AuthRepositoryImpl struct {
    // No need for fields if DB is passed in methods
}

func NewAuthRepository() AuthRepository {
    return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, db *gorm.DB, auth domain.Auth) domain.Auth {
	return auth
}