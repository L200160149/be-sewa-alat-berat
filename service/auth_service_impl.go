package service

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/L200160149/be-sewa-alat-berat/exception"
	"github.com/L200160149/be-sewa-alat-berat/helper"
	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"github.com/L200160149/be-sewa-alat-berat/model/web"
	"github.com/L200160149/be-sewa-alat-berat/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, db *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request web.AuthRequest) web.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	var user domain.User
	err = service.DB.WithContext(ctx).Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		panic(exception.NewUnauthorizedError("Invalid email or password"))
	}

	// Verify password
	if err := helper.ComparePassword(user.Password, request.Password); err != nil {
		panic(exception.NewUnauthorizedError("Invalid email or password"))
	}

	
	// Create JWT token
	secret := os.Getenv("JWT_SECRET")
	expiredStr := os.Getenv("JWT_EXPIRED")
	expired, _ := strconv.Atoi(expiredStr)

	claims := jwt.MapClaims{
		"email": user.Email,
		"name": user.Name,
		"role": user.Role,
		"exp": time.Now().Add(time.Duration(expired) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	helper.PanicIfError(err)

	return web.AuthResponse{AccessToken: tokenString}
}
