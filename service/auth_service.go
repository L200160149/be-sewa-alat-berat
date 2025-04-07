package service

import (
	"context"

	"github.com/L200160149/be-sewa-alat-berat/model/web"
)

type AuthService interface {
	Login(ctx context.Context, request web.AuthRequest) web.AuthResponse
}