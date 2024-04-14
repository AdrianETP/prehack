package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	app := fiber.New()
	// conexion a la base de datos
	MongoConnect()

	// llamadas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": 200,
			"result": "Hello world",
		})
	})

	app.Post("/AddUser", func(c *fiber.Ctx) error {
		// obtener el usuario del body
		var user User

		err := c.BodyParser(&user)
		if err != nil {
			log.Fatal("Error decoding user")
		}
		user.Id = primitive.NewObjectID()
		res, err := Database.Collection("Users").InsertOne(context.TODO(), user)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": 400,
				"result": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"status": 200,
			"result": res,
		})
	})

	app.Post("/Login", func(c *fiber.Ctx) error {
		var user User
		err := c.BodyParser(&user)
		if err != nil {
			log.Fatal("Error decoding body")
		}

		res := Database.Collection("Users").FindOne(context.TODO(), bson.D{{"username", user.UserName}})
		var LogedUser User
		err = res.Decode(&LogedUser)
		if err != nil {
			log.Fatal("Error decoding logedUser")
		}

		if LogedUser.Password != user.Password {
			return c.Status(400).JSON(fiber.Map{
				"status": 400,
				"error":  "usario y/o contraseña incorrecta",
			})
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		return c.Status(200).JSON(fiber.Map{
			"status": 200,
			"result": user,
		})
	})

	// creacion de la app
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error listening on port 3000")
	}
}
