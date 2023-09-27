package routes

import (
	"github.com/MatkoMilic/GO-fiber-gorm/database"
	"github.com/MatkoMilic/GO-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	// the serializer
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

type UpdateProductStr struct {
	// the update type
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}

	if err := c.BodyParser(&product); err != nil {
		c.Status(422).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProduct(c *fiber.Ctx) error {
	product := models.Product{}

	productId, err := c.ParamsInt("id")

	if err != nil {
		c.Status(422).JSON(err.Error())
	}

	database.Database.Db.Find(&product, productId)

	if product.ID == 0 {
		return c.Status(400).JSON("User for given ID was not found!")
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	modelProducts := []models.Product{}

	database.Database.Db.Find(&modelProducts)

	products := []Product{}

	for _, modelProduct := range modelProducts {
		responseProduct := CreateResponseProduct(modelProduct)
		products = append(products, responseProduct)
	}

	return c.Status(200).JSON(products)
}

func DeleteProduct(c *fiber.Ctx) error {
	product := models.Product{}

	productId, err := c.ParamsInt("id")

	if err != nil {
		c.Status(422).JSON(err.Error())
	}

	database.Database.Db.Delete(&product, productId)

	return c.Status(200).JSON("Successfully deleted the product.")
}

func UpdateProduct(c *fiber.Ctx) error {
	product := models.Product{}
	updateProduct := UpdateProductStr{}

	productId, err := c.ParamsInt("id")

	if err != nil {
		c.Status(422).JSON(err.Error())
	}

	parsingError := c.BodyParser(&updateProduct)

	if parsingError != nil {
		c.Status(422).JSON("Arguments are not in the proper format.")
	}

	database.Database.Db.Find(&product, productId)

	if product.ID == 0 {
		c.Status(404).JSON("Product with given ID is not found!")
	}

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}
