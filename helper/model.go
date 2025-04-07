package helper

import (
	"github.com/L200160149/be-sewa-alat-berat/model/domain"
	"github.com/L200160149/be-sewa-alat-berat/model/web"
)

func ToUsersResponse(user domain.User) web.UsersResponse {
    return web.UsersResponse{
        Id:    user.Id,
        Name:  user.Name,
        Email: user.Email,
        Role:  user.Role,
    }
}

func ToUsersResponses(categories []domain.User) []web.UsersResponse {
	var usersResponses []web.UsersResponse
	for _, users := range categories {
		usersResponses = append(usersResponses, ToUsersResponse(users))
	}
	return usersResponses
}
