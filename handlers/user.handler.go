package handlers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/models"
	m "github.com/haji-sudo/ShabehRoshan/models/validation"
	"github.com/haji-sudo/ShabehRoshan/repository"
	"github.com/haji-sudo/ShabehRoshan/util"
)

func Profile(c *fiber.Ctx) error {
	userid := c.Locals("userid").(string)
	repo := repository.NewUserRepository()
	user, err := repo.GetByID(uuid.MustParse(userid))
	repo.GetProfile(user)
	if err != nil {
		return nil
	}

	return c.Render("user/profile", fiber.Map{"user": user}, Layout)
}

func Settings(c *fiber.Ctx) error {
	userid := c.Locals("userid").(string)
	repo := repository.NewUserRepository()
	user, err := repo.GetByID(uuid.MustParse(userid))
	repo.GetProfile(user)
	if err != nil {
		return nil
	}

	return c.Render("user/settings", fiber.Map{"user": user}, Layout)
}
func UpdateProfile(c *fiber.Ctx) error {
	name := c.FormValue("name")
	username := c.FormValue("username")
	bio := c.FormValue("bio")
	photo, loadPhoto := c.FormFile("profilePicture")

	data := models.User{Username: username, Profile: models.Profile{Name: name, Bio: bio}}
	errorData := new(m.UpdateProfile)
	repo := repository.NewUserRepository()
	//Check Input IsValid
	err := util.ValidateUpdateProfileInput(name, username, bio)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			if v.Field() == "Username" {
				errorData.Username = "Username is not valid !"
			} else if v.Field() == "Name" {
				errorData.Name = "Name is not valid !"
			} else if v.Field() == "Bio" {
				errorData.Bio = "Max length 256"
			}
		}
		return c.Render("user/settings", fiber.Map{
			"error": errorData,
			"user":  data})
	} else if loadPhoto == nil {
		if photo.Size > (2 * 1024 * 1024) {
			errorData.Photo = "Max size 2MB"
			return c.Render("user/settings", fiber.Map{
				"error": errorData,
				"user":  data})
		}
	}
	userid := c.Locals("userid").(string)
	user, err := repo.GetByID(uuid.MustParse(userid))
	if err != nil {
		return c.Redirect("/")
	}
	if user.Username != strings.ToLower(username) {
		if ok, _ := repo.UsernameExist(username); ok {
			errorData.Username = "Username is Exist !"
			return c.Render("user/settings", fiber.Map{
				"error": errorData,
				"user":  data})
		}
	}
	repo.GetProfile(user)
	if loadPhoto == nil {
		photoName, err := util.SavePhotoAndOptimze(photo)
		if err != nil {
			errorData.Photo = "Something wrong in saving photo"
			return c.Render("user/settings", fiber.Map{
				"error": errorData,
				"user":  data})
		}
		user.Profile.Picture = photoName
	}

	user.Profile.Name = name
	user.Profile.Bio = bio
	user.Username = strings.ToLower(username)

	repo.Update(user)
	repo.UpdateProfile(user)

	return c.Redirect("/Dashboard/profile")

}
