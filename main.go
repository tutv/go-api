package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tutv/go-api/database"
	"github.com/tutv/go-api/routes"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome api.")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)

	app.Post("/users", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/users/:id", routes.GetUser)
	app.Put("/users/:id", routes.UpdateUser)
	app.Delete("/users/:id", routes.DeleteUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
