package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/repository"
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
