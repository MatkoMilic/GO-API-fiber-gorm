package main

import (
	"log"

	"github.com/MatkoMilic/GO-fiber-gorm/database"
	"github.com/MatkoMilic/GO-fiber-gorm/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Matko's GO API!")
}

func setupRoutes(app *fiber.App) {
	//get welcome
	app.Get("/api", welcome)
	//post user
	app.Post("/api/users", routes.CreateUser)
	//get users
	app.Get("/api/get-users", routes.GetUsers)
	//get user
	app.Get("/api/get-user/:id", routes.GetUser)
	//update user
	app.Put("/api/update-user/:id", routes.UpdateUser)
	//delete user
	app.Delete("/api/delete-user/:id", routes.DeleteUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
