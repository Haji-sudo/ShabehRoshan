package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/haji-sudo/ShabehRoshan/handlers"
	"github.com/haji-sudo/ShabehRoshan/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup routes
	app.Get(Home, handlers.Home)
	app.Get(About, handlers.About)
	app.Get(Test, handlers.Test)

	app.Get(SignUp, handlers.SignUp)
	app.Get(Login, handlers.Login)
	app.Get(VerifyEmail, handlers.VerifyEmail)
	app.Get(ResendVerifyEmail, handlers.ResendEmail)
	//
	app.Post(SignUp, handlers.SignupUser)
	app.Post(Login, handlers.LoginUser)
	app.Post(ResendVerifyEmail, handlers.ResendVerifyEmail)

	protected := app.Group(UserPanel, middleware.Auth)
	protected.Get(Logout, handlers.LogOut)
	protected.Get(About, handlers.About)
	protected.Get(Profile, handlers.Profile)
	protected.Get(Settings, handlers.Settings)
	protected.Post(Settings, handlers.UpdateProfile)
	// Setup static files
	app.Static("/public", "./public")

	// Handle not founds
	app.Use(handlers.NotFound)
	protected.Use(handlers.NotFound)

}
