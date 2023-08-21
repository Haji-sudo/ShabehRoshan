package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/haji-sudo/ShabehRoshan/handlers"
	"github.com/haji-sudo/ShabehRoshan/middleware"
	"github.com/haji-sudo/ShabehRoshan/router/url"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup url
	app.Get(url.Home, handlers.Home)
	app.Get(url.About, handlers.About)
	app.Get(url.Test, handlers.Test)

	app.Get(url.SignUp, handlers.SignUp)
	app.Get(url.Login, handlers.Login)
	app.Get(url.VerifyEmail, handlers.VerifyEmail)
	app.Get(url.ResendVerifyEmail, handlers.ResendVerifyEmail)
	app.Get(url.ForgotPassword, handlers.ForgotPassword)
	app.Get(url.ResetPassword, handlers.ResetPassword)
	//
	app.Post(url.SignUp, handlers.SignUp)
	app.Post(url.Login, handlers.Login)
	app.Post(url.ResendVerifyEmail, handlers.ResendVerifyEmail)
	app.Post(url.ForgotPassword, handlers.ForgotPassword)
	app.Post(url.ResetPassword, handlers.ResetPassword)

	app.Get(url.Post+"/:postID", handlers.GetPost)

	protected := app.Group(url.UserPanel, middleware.Auth)
	protected.Get(url.Logout, handlers.LogOut)
	protected.Get(url.About, handlers.About)
	protected.Get(url.Profile, handlers.Profile)
	protected.Get(url.Settings, handlers.Settings)
	protected.Post(url.Settings, handlers.UpdateProfile)
	protected.Get(url.CreatePost, handlers.CreatePost)
	protected.Post(url.CreatePost, handlers.CreatePost)
	// Setup static files
	app.Static(url.Static, "./public")

	// Handle not founds
	app.Use(handlers.NotFound)
	protected.Use(handlers.NotFound)

}
