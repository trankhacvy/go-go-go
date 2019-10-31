package users

type LoginValidator struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterValidator struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
}