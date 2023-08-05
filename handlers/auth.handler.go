package handlers

import (
	"log"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/db"
	"github.com/haji-sudo/ShabehRoshan/middleware"
	"github.com/haji-sudo/ShabehRoshan/models"
	v "github.com/haji-sudo/ShabehRoshan/models/validation"
	"github.com/haji-sudo/ShabehRoshan/repository"
	"github.com/haji-sudo/ShabehRoshan/services"
	"github.com/haji-sudo/ShabehRoshan/util"
)

func SignUp(c *fiber.Ctx) error {
	if !middleware.CheckAuthNotValid(c) {
		return c.Redirect("/")
	}
	return c.Render("user/signup", nil, Layout)
}

func SignupUser(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	data := v.SignUpUser{Username: username, Email: email, Password: ""}
	//Check Input IsValid
	err := util.ValidateSignupInput(username, email, password)
	if err != nil {
		errorData := new(v.SignUpUser)
		for _, v := range err.(validator.ValidationErrors) {
			if v.Field() == "Email" {
				errorData.Email = "Email is not valid !"
				data.Email = ""
			} else if v.Field() == "Username" {
				errorData.Username = "Username is not valid !"
				data.Email = ""
			} else if v.Field() == "Password" {
				errorData.Password = "The minimum length is 8 characters !"
			}
		}
		return c.Render("user/signup", fiber.Map{
			"error": errorData,
			"data":  data})
	}

	//Check Input Not Exist
	repo := repository.NewUserRepository()
	chEmail, _ := repo.EmailExist(email)
	chUsername, _ := repo.UsernameExist(username)
	errData := new(v.SignUpUser)
	if chEmail || chUsername {
		if chEmail {
			errData.Email = "Email is exist"
			data.Email = ""
		} else if chUsername {
			errData.Username = "Username is exist"
			data.Username = ""
		}
		return c.Render("user/signup", fiber.Map{
			"error": errData,
			"data":  data,
		})
	}

	// All Validation Pass Now Create User
	hashpw, err := util.HashPassword(password)
	if err != nil {
		return c.Render("user/error", "Something Wrong in Password hashing try again !!")
	}

	user := models.User{ID: uuid.New(), Username: strings.ToLower(username), Email: strings.ToLower(email), PasswordHash: string(hashpw)}
	err = repo.Create(&user)
	if err != nil {
		repo.Delete(&user)
		return c.Render("user/error", "Something Wrong in create user try again !!")
	}
	emailValidateToken, err := util.GenerateValidationToken(user)
	if err != nil {
		repo.Delete(&user)
		return c.Render("user/error", "Failed to generate token")
	}
	repo.GetToken(&user)
	user.Token.VerifyEmailToken = emailValidateToken
	err = services.SendingVerificationEmail(user.Email, user.Username, emailValidateToken)
	if err != nil {
		repo.Delete(&user)
		return c.Render("user/error", "Failed to Verification Email")
	}
	repo.Update(&user)
	return c.Render("user/successSignup", data)
}

func Login(c *fiber.Ctx) error {
	if !middleware.CheckAuthNotValid(c) {
		return c.Redirect("/")
	}
	return c.Render("user/login", nil, Layout)
}

func LoginUser(c *fiber.Ctx) error {
	middleware.CheckAuthNotValid(c)
	email := c.FormValue("email")
	password := c.FormValue("password")
	remember_me := c.FormValue("remember_me")
	data := v.LoginUser{Email: email, Password: ""}
	err := util.ValidateLoginInput(email, password)
	if err != nil {
		errorData := new(v.LoginUser)
		for _, v := range err.(validator.ValidationErrors) {
			if v.Field() == "Email" {
				errorData.Email = "Email is not valid !"
				data.Email = ""
			} else if v.Field() == "Password" {
				errorData.Password = "The minimum length is 8 characters !"
			}
		}
		return c.Status(200).Render("user/login", fiber.Map{
			"error": errorData,
			"data":  data,
		})
	}

	repo := repository.NewUserRepository()
	//Check Email Exist
	if u, _ := repo.EmailExist(email); !u {
		return c.Status(200).Render("user/login", fiber.Map{
			"error": v.LoginUser{
				Email: "Couldn’t find a account associated with this email. Try again or create an account .",
			},
		})
	}

	if user, err := repo.GetByEmail(email); err != nil {
		log.Println(err)
		return c.Status(200).Render("user/login", fiber.Map{"NotSuccess": "Something went wrong."})
	} else if err := util.ComparePassword([]byte(user.PasswordHash), password); err != nil {
		return c.Status(200).Render("user/login", fiber.Map{"NotSuccess": "Login unsuccessful. Please check your credentials and try again."})
	} else if !user.EmailVerified {
		return c.Status(200).Render("user/login", fiber.Map{"NotSuccess": "Your email has not been verified yet. Confirm your email first", "notverfied": user.Email})
	} else if !user.Active {
		c.Status(200).Render("user/login", fiber.Map{"NotSuccess": "Your account has been deactivated"})
	} else {
		user.LastLogin = time.Now()
		repo.Update(user)
		token, err := util.GenerateToken(*user)
		if err != nil {
			return c.Status(200).Render("user/login", fiber.Map{"NotSuccess": "Failed to generate token"})
		}

		//Check if remember me checked create a RefreshToken
		repo.GetToken(user)
		user.Token.RefreshToken = ""
		if remember_me == "on" {
			refreshToken, err := util.GenerateRefreshToken(*user)
			if err != nil {
				return c.Status(200).Render("user/login", fiber.Map{"NotSuccess": "Failed to generate refresh token"})
			}
			user.Token.RefreshToken = refreshToken
		}
		repo.UpdateToken(user)
		sess, _ := db.Store.Get(c)
		sess.Set("token", token)
		sess.Save()
		return c.Status(200).Render("user/successLogin", v.SignUpUser{Username: user.Username, Email: user.Email})
	}

	return nil
}

func LogOut(c *fiber.Ctx) error {
	sess, err := db.Store.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}
	sess.Destroy()

	return c.Redirect("/")
}

func VerifyEmail(c *fiber.Ctx) error {
	token := c.Query("token")

	if token == "" {
		return c.Render("user/error", "Missing verification token")
	}
	userid, err := util.ValidateToken(token)
	if err != nil {
		return c.Render("user/error", "The token is not valid")
	}
	repo := repository.NewUserRepository()

	user, err := repo.GetByID(uuid.MustParse(userid))
	if err != nil {
		return c.Render("user/error", "The token is not valid")
	}

	repo.GetToken(user)
	if user.Token.VerifyEmailToken != token {
		return c.Render("user/error", "The token is not valid")
	}

	user.EmailVerified = true
	user.Active = true
	user.Token.VerifyEmailToken = ""
	repo.Update(user)
	repo.UpdateToken(user)
	return c.Render("user/verification-success", nil)
}
func ResendEmail(c *fiber.Ctx) error {
	middleware.CheckAuthNotValid(c)
	return c.Render("user/ResendEmail", nil, Layout)

}
func ResendVerifyEmail(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if email == "" {
		err := c.Render("user/ResendEmail", fiber.Map{"error": "Email is required"})
		return err
	}

	err := util.ValidateEmail(email)
	if err != nil {
		return c.Render("user/ResendEmail", fiber.Map{"error": "Email is not valid"})
	}

	repo := repository.NewUserRepository()

	// Check if email exists
	if exists, _ := repo.EmailExist(email); !exists {
		return c.Render("user/ResendEmail", fiber.Map{"error": "Couldn't find an account associated with this email. Try again or create an account."})
	}

	user, _ := repo.GetByEmail(email)

	if user.EmailVerified {
		return c.Render("user/ResendEmail", fiber.Map{"NotSuccess": "This email has been verified"})
	}

	validateToken, err := util.GenerateValidationToken(*user)
	if err != nil {
		return c.Render("user/ResendEmail", fiber.Map{"NotSuccess": "Failed to generate token"})
	}

	err = services.SendingVerificationEmail(user.Email, user.Username, validateToken)
	if err != nil {
		return c.Render("user/ResendEmail", fiber.Map{"NotSuccess": "Failed to send email. Please try again."})
	}
	repo.GetToken(user)
	user.Token.VerifyEmailToken = validateToken
	repo.UpdateToken(user)

	return c.Render("user/ResendEmail", fiber.Map{"Success": "Confirmation email has been sent"})
}
