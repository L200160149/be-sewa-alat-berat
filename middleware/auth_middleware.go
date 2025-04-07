package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/L200160149/be-sewa-alat-berat/helper"
	"github.com/L200160149/be-sewa-alat-berat/model/web"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/api/v1/login" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	authHeader := request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		unauthorized(writer)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		unauthorized(writer)
		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}

func unauthorized(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
