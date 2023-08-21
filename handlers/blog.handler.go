package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/models"
	m "github.com/haji-sudo/ShabehRoshan/models/validation"
	"github.com/haji-sudo/ShabehRoshan/repository"
	"github.com/haji-sudo/ShabehRoshan/util"
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
	photo, loadPhoto := c.FormFile("coverimage")
	postData := m.CreatePost{Title: c.FormValue("title"), Content: c.FormValue("content"), Tag: c.FormValue("tags")}
	errorData := new(m.CreatePost)
	err := util.ValidateCreatePost(postData)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			if v.Field() == "Title" {
				errorData.Title = "Max Character is 100"
			} else if v.Field() == "Content" {
				errorData.Content = "Max Character is 10,000"
			} else if v.Field() == "Tag" {
				errorData.Tag = "Max Character is 100"
			}
		}
		return c.Render("blog/createblog", fiber.Map{"error": errorData, "data": postData})
	} else if loadPhoto == nil {
		if photo.Size > (5 * 1024 * 1024) {
			errorData.Photo = "Max size is 5MB"
			return c.Render("blog/createblog", fiber.Map{"error": errorData, "data": postData})
		}
	}

	userid := c.Locals("userid").(string)
	urepo := repository.NewUserRepository()
	user, _ := urepo.GetByID(uuid.MustParse(userid))
	imgname, _ := util.SaveImageAndOptimize(photo)
	brepo := repository.NewBlogRepository()
	post := models.Post{ID: uuid.New(), Title: postData.Title, Content: postData.Content, CoverImage: imgname, User: *user}
	brepo.Create(&post)
	return c.SendString("Done")
}

func GetPost(c *fiber.Ctx) error {
	postID, err := uuid.Parse(c.Params("postID"))
	if err != nil {
		return c.SendString("Id Not Valid")
	}
	repo := repository.NewBlogRepository()
	post, err := repo.GetByID(postID)
	if err != nil {
		return c.SendString("Post not found")
	}
	return c.Render("blog/blogArticle", fiber.Map{"post": post}, Layout)

}
