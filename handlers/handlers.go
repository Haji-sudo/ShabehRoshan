package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/db"
	"github.com/haji-sudo/ShabehRoshan/middleware"
	"github.com/haji-sudo/ShabehRoshan/repository"
	"github.com/haji-sudo/ShabehRoshan/util"
)

const Layout = "layouts/main"

// Home renders the home view
func Home(c *fiber.Ctx) error {
	if middleware.IsAuth(c) {
		repo := repository.NewUserRepository()
		userid := c.Locals("userid").(string)
		user, _ := repo.GetByID(uuid.MustParse(userid))
		repo.GetProfile(user)
		return c.Render("index", fiber.Map{
			"Title": "Hello", "user": user,
		}, Layout)

	}
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	}, Layout)
}

// About renders the about view
func About(c *fiber.Ctx) error {
	return c.Render("about", nil)
}

func Test(c *fiber.Ctx) error {
	session_id := c.Cookies("session_id")
	if session_id == "" {
		return c.Redirect("/login")
	}
	sess, err := db.Store.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}
	var tokenString string
	if token, ok := sess.Get("token").(string); ok {
		tokenString = token
	}
	if tokenString == "" {
		sess.Destroy()
		return c.Redirect("/login")
	}
	userid, _ := util.ValidateToken(tokenString)

	repo := repository.NewUserRepository()
	user, _ := repo.GetByID(uuid.MustParse(userid))
	fmt.Println(user)
	// fmt.Println("|||||||||||||||||||||||||||||||||\n\n")
	repo.GetToken(user)

	fmt.Println(user)
	return nil
}

// NoutFound renders the 404 view
func NotFound(c *fiber.Ctx) error {
	if middleware.IsAuth(c) {
		return c.Status(404).Render("partials/404", fiber.Map{"user": "test"}, Layout)
	}
	return c.Status(404).Render("partials/404", nil, Layout)
}
