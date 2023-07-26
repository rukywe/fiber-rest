package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rukywe/fiber-rest/database"
	"github.com/rukywe/fiber-rest/models"
)

type User struct {
	//This is not the model user
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User ) User {

	return User{
		ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName,
	}
}


func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}


func GetUsers(c *fiber.Ctx) error {
		users := []models.User{}

		database.Database.Db.Find(&users)
		responseUsers := []User{}

		for _, user := range users {
			responseUser := CreateResponseUser(user)
			responseUsers = append(responseUsers, responseUser)
		}

		 	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User)error{

	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("user not found")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id , err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure ID is an Integer")
	}

	if err:= findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}


func UpdateUser(c *fiber.Ctx) error {
	id , err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure ID is an Integer")
	}

	if err:= findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil{
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error {

	id , err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure ID is an Integer")
	}

	if err:= findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil{
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted user")

}