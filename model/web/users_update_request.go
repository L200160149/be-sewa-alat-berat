package web

type UsersUpdateRequest struct {
	Id       int    `json:"id"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
	Role     string `validate:"required" json:"role"`
}
