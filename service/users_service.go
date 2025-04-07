package service

import (
	"context"

	"github.com/L200160149/be-sewa-alat-berat/model/web"
)

type UsersService interface {
	FindAll(ctx context.Context) []web.UsersResponse
	Create(ctx context.Context, request web.UsersCreateRequest) web.UsersResponse
	Update(ctx context.Context, request web.UsersUpdateRequest) web.UsersResponse
	Delete(ctx context.Context, userId int) web.UsersResponse
	
}