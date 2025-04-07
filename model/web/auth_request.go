package web

type AuthRequest struct {
	Email		string	`validate:"required,min=1,max=100" json:"email"`
	Password	string	`validate:"required,min=8" json:"password"`
}