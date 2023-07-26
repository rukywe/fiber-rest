package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rukywe/fiber-rest/database"
	"github.com/rukywe/fiber-rest/models"
)


type Product struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{productModel.ID, productModel.Name, productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error{
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}