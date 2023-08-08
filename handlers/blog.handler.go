package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/models"
	"github.com/haji-sudo/ShabehRoshan/repository"
)

func CreatePost(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		userid := c.Locals("userid").(string)
		repo := repository.NewUserRepository()
		user, err := repo.GetByID(uuid.MustParse(userid))
		repo.GetProfile(user)
		if err != nil {
			return nil
		}
		return c.Render("blog/createblog", fiber.Map{"user": user}, Layout)
	}
	blog := new(models.Blog)
	blog.AuthorID = uuid.MustParse(c.Locals("userid").(string))
	blog.ID = uuid.New()
	blog.PublishDate = time.Now()
	blog.Title = c.FormValue("title")
	blog.Content = c.FormValue("content")
	repo := repository.NewBlogRepository()

	repo.Create(blog)
	return c.SendString("Done")
}
