package validators

//import "gopkg.in/go-playground/validator.v9"

type LoginValidator struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}