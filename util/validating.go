package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
	vm "github.com/haji-sudo/ShabehRoshan/models/validation"
)

func ValidateEmail(email string) error {
	validate := validator.New()
	if err := validate.Var(strings.ToLower(email), "required,email"); err != nil {
		return err
	}
	return nil
}
func ValidateUsername(username string) error {
	validate := validator.New()
	if err := validate.Var(strings.ToLower(username), "required,min=4,max=20,alphanum"); err != err {
		return err
	}
	return nil
}
func ValidateSignupInput(username, email, password string) error {
	user := vm.SignUpUser{
		Username: strings.ToLower(username),
		Email:    strings.ToLower(email),
		Password: password,
	}
	// Create a new validator instance
	validate := validator.New()
	// Validate the struct fields
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
func ValidateLoginInput(email, password string) error {
	user := vm.LoginUser{
		Email:    strings.ToLower(email),
		Password: password,
	}
	// Create a new validator instance
	validate := validator.New()
	// Validate the struct fields
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
