package repository

import (
	"context"

	"github.com/L200160149/be-sewa-alat-berat/exception"
	"github.com/L200160149/be-sewa-alat-berat/helper"
	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"gorm.io/gorm"
)
type UsersRepositoryImpl struct {
    // No need for fields if DB is passed in methods
}

func NewUsersRepository() UsersRepository {
    return &UsersRepositoryImpl{}
}

func (repository *UsersRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) []domain.User {
    var users []domain.User
    err := db.WithContext(ctx).Table("users").Find(&users).Error
    helper.PanicIfError(err)
    return users
}

func (repository *UsersRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	err := db.WithContext(ctx).Table("users").Create(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UsersRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	result := db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", user.Id).Updates(user)

	if result.RowsAffected == 0 {
		panic(exception.NewNotFoundError("User not found"))
	}
	
	helper.PanicIfError(result.Error)
	return user
}

func (repository *UsersRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	result := db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", user.Id).Delete(user)

	if result.RowsAffected == 0 {
		panic(exception.NewNotFoundError("User not found"))
	}
	
	helper.PanicIfError(result.Error)
	return user
}
