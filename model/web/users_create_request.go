package web

type UsersCreateRequest struct {
	Name		string	`validate:"required,min=1,max=100" json:"name"`
	Email		string	`validate:"required,min=1,max=100" json:"email"`
	Password	string	`validate:"required,min=8" json:"password"`
	Role		string	`validate:"required" json:"role"`
}