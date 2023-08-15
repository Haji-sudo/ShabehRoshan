package models

type SignUpUser struct {
	Name     string `validate:"required,min=2,max=20"`
	Username string `validate:"required,min=4,max=20,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type LoginUser struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type UpdateProfile struct {
	Photo    string
	Name     string `validate:"required,min=2,max=20"`
	Username string `validate:"required,min=4,max=20,alphanum"`
	Bio      string `validate:"max=256"`
}

type CreatePost struct {
	Title   string `validate:"max=100"`
	Content string `validate:"max=10000"`
	Tag     string `validate:"max=100"`
	Photo   string
}
