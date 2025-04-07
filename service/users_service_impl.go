package service

import (
	"context"
	"strings"

	"github.com/L200160149/be-sewa-alat-berat/exception"
	"github.com/L200160149/be-sewa-alat-berat/helper"
	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"github.com/L200160149/be-sewa-alat-berat/model/web"
	"github.com/L200160149/be-sewa-alat-berat/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersServiceImpl struct {
    UsersRepository repository.UsersRepository
    DB              *gorm.DB
    Validate        *validator.Validate
}

func NewUsersService(usersRepository repository.UsersRepository, db *gorm.DB, validate *validator.Validate) UsersService {
    return &UsersServiceImpl{
        UsersRepository: usersRepository,
        DB:              db,
        Validate:        validate,
    }
}

func (service *UsersServiceImpl) FindAll(ctx context.Context) []web.UsersResponse {
    var users []domain.User
    err := service.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        users = service.UsersRepository.FindAll(ctx, tx)
        return nil
    })
    helper.PanicIfError(err)
    return helper.ToUsersResponses(users)
}

func (service *UsersServiceImpl) Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	var response web.UsersResponse

	err = service.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		helper.PanicIfError(err)

		user := domain.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashedPassword),
			Role:     request.Role,
		}

		// catch duplicate email here
		defer func() {
			if r := recover(); r != nil {
				if errMsg, ok := r.(error); ok && strings.Contains(errMsg.Error(), "Error 1062") {
					panic(exception.NewBadRequestError("Email already exists"))
				}
				panic(r)
			}
		}()

		user = service.UsersRepository.Save(ctx, tx, user)
		response = helper.ToUsersResponse(user)
		return nil
	})

	helper.PanicIfError(err)
	return response
}

func (service *UsersServiceImpl) Update(ctx context.Context, request web.UsersUpdateRequest) web.UsersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	var response web.UsersResponse

	err = service.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		helper.PanicIfError(err)

		user := domain.User{
			Id:     request.Id,
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashedPassword),
			Role:     request.Role,
		}

		// catch duplicate email here
		defer func() {
			if r := recover(); r != nil {
				if errMsg, ok := r.(error); ok && strings.Contains(errMsg.Error(), "Error 1062") {
					panic(exception.NewBadRequestError("Email already exists"))
				}
				panic(r)
			}
		}()

		user = service.UsersRepository.Update(ctx, tx, user)
		response = helper.ToUsersResponse(user)
		return nil
	})

	helper.PanicIfError(err)
	return response
}

func (service *UsersServiceImpl) Delete(ctx context.Context, userId int) web.UsersResponse {
	var response web.UsersResponse

	err := service.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := domain.User{
			Id: userId,
		}

		user = service.UsersRepository.Delete(ctx, tx, user)
		response = helper.ToUsersResponse(user)
		return nil
	})

	helper.PanicIfError(err)
	return response
}
