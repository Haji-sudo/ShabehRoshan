package main

import (
	"github.com/haji-sudo/ShabehRoshan/config"
	"github.com/haji-sudo/ShabehRoshan/db"
	"github.com/haji-sudo/ShabehRoshan/router"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//Init Configures
	config.Init()

	// Init Database
	db.Init()
	db.InitSession()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Views: config.Engine,
	})
	app.Use(cors.New())
	// app.Use(csrf.New())
	//Setup Routes
	router.SetupRoutes(app)

	//Start Web Server
	log.Fatal(app.Listen(":3000"))
}
