package repository

import (
	"context"

	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"gorm.io/gorm"
)

type UsersRepository interface {
    FindAll(ctx context.Context, db *gorm.DB) []domain.User
	Save(ctx context.Context, db *gorm.DB, user domain.User) domain.User
	Update(ctx context.Context, db *gorm.DB, user domain.User) domain.User
	Delete(ctx context.Context, db *gorm.DB, user domain.User) domain.User
}