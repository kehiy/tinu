package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kehiy/tinu/model"
)

func getAll(c *fiber.Ctx) error {
	tinus, err := model.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"internal server error" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":tinus,
	})
}

func getOne(c *fiber.Ctx) error {
	id := c.Params("id")
	tinu, err := model.GetOne(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":tinu,
	})
}


func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/tinu", getAll)
	router.Get("/tinu/:id", getOne)

	router.Listen(":3000")
}