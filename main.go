package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rukywe/fiber-rest/database"
	"github.com/rukywe/fiber-rest/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App){

	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	//Products

	app.Post("/api/products", routes.CreateProduct)



}

func main() {

	database.ConnectDb()
	app:=fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}