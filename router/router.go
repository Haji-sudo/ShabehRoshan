package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/haji-sudo/ShabehRoshan/handlers"
	"github.com/haji-sudo/ShabehRoshan/middleware"
	"github.com/haji-sudo/ShabehRoshan/router/routes"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup routes
	app.Get(routes.Home, handlers.Home)
	app.Get(routes.About, handlers.About)
	app.Get(routes.Test, handlers.Test)

	app.Get(routes.SignUp, handlers.SignUp)
	app.Get(routes.Login, handlers.Login)
	app.Get(routes.VerifyEmail, handlers.VerifyEmail)
	app.Get(routes.ResendVerifyEmail, handlers.ResendVerifyEmail)
	app.Get(routes.ForgotPassword, handlers.ForgotPassword)
	app.Get(routes.ResetPassword, handlers.ResetPassword)
	//
	app.Post(routes.SignUp, handlers.SignUp)
	app.Post(routes.Login, handlers.Login)
	app.Post(routes.ResendVerifyEmail, handlers.ResendVerifyEmail)
	app.Post(routes.ForgotPassword, handlers.ForgotPassword)
	app.Post(routes.ResetPassword, handlers.ResetPassword)

	protected := app.Group(routes.UserPanel, middleware.Auth)
	protected.Get(routes.Logout, handlers.LogOut)
	protected.Get(routes.About, handlers.About)
	protected.Get(routes.Profile, handlers.Profile)
	protected.Get(routes.Settings, handlers.Settings)
	protected.Post(routes.Settings, handlers.UpdateProfile)
	// Setup static files
	app.Static("/public", "./public")

	// Handle not founds
	app.Use(handlers.NotFound)
	protected.Use(handlers.NotFound)

}
