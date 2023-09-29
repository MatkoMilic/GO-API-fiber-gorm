package routes

import (
	"errors"
	"time"

	"github.com/MatkoMilic/GO-fiber-gorm/database"
	"github.com/MatkoMilic/GO-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	// the serializer
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	order := models.Order{}

	if err := c.BodyParser(&order); err != nil {
		c.Status(400).JSON(err.Error())
	}

	user := models.User{}

	database.Database.Db.Find(&user, order.UserRefer)

	if user == (models.User{}) {
		c.Status(400).SendString("User does not exist!")
	}

	product := models.Product{}

	database.Database.Db.Find(&user, order.ProductRefer)

	if product == (models.Product{}) {
		c.Status(400).SendString("Product does not exist!")
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		user := models.User{}
		product := models.Product{}
		database.Database.Db.Find(&user, order.UserRefer)
		database.Database.Db.Find(&product, order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, id)
	if order.ID == 0 {
		return errors.New("Order does not exist.")
	}

	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	order := models.Order{}

	if err != nil {
		return c.Status(400).JSON("ID is not valid!")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&order, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}
