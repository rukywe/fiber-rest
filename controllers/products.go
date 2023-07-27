package controllers

import (
	"errors"

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

	return c.Status(201).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := []Product{}

	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

		 return c.Status(200).JSON(responseProducts)
}


func findProduct(id int, product *models.Product)error{

	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("Product not found")
	}

	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id , err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure ID is an Integer")
	}

	if err:= findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}


func UpdateProduct(c *fiber.Ctx) error {
	id , err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure ID is an Integer")
	}

	if err:= findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil{
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}


func DeleteProduct(c *fiber.Ctx) error {

	id , err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure ID is an Integer")
	}

	if err:= findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil{
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted product")

}

