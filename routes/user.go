package routes

import (
	"github.com/MatkoMilic/GO-fiber-gorm/database"
	"github.com/MatkoMilic/GO-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	// the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUserStr struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
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

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	user := models.User{}

	if err != nil {
		return c.Status(400).JSON("ID is not valid!")
	}

	database.Database.Db.First(&user, id)

	if user.ID == 0 {
		return c.Status(400).JSON("User for given ID was not found!")
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	user := models.User{}

	if err != nil {
		return c.Status(400).JSON("ID is not valid!")
	}

	var updateUserData UpdateUserStr

	if err := c.BodyParser(&updateUserData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.First(&user, id)

	if user.ID == 0 {
		return c.Status(400).JSON("User for given ID was not found!")
	}

	user.FirstName = updateUserData.FirstName
	user.LastName = updateUserData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	user := models.User{}

	if err != nil {
		return c.Status(400).JSON("ID is not valid!")
	}

	database.Database.Db.First(&user, id)

	if user.ID == 0 {
		return c.Status(400).JSON("User for given ID was not found!")
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully deleted.")
}
