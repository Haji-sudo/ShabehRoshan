package models

type SignUpUser struct {
	Username string `validate:"required,min=4,max=20,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type LoginUser struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
