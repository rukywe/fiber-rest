package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rukywe/fiber-rest/controllers"
	"github.com/rukywe/fiber-rest/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App){

	app.Get("/api", welcome)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	//Products
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products", controllers.GetProducts)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	// orders

	app.Post("/api/orders", controllers.CreateOrder)
	app.Get("/api/orders", controllers.GetOrders)



}

func main() {

	database.ConnectDb()
	app:=fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}