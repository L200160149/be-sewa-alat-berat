package repository

import (
	"context"

	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(ctx context.Context, db *gorm.DB, auth domain.Auth) domain.Auth
}