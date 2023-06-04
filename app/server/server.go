package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kehiy/tinu/model"
	"github.com/kehiy/tinu/utils"
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

func createTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var tinu model.Tinu
	err := c.BodyParser(&tinu)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	tinu.ID = utils.GenerateId()
	err = model.CreateTinu(tinu)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"OK",
		"data":tinu,
	})
}

func tinu(c *fiber.Ctx) error {
	id := c.Params("id")
	tinu, err := model.GetOne(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"internal server error",
		})
	}
	tinu.Clicked += 1
	err = model.UpdateTinu(tinu)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"internal server error",
		})
	}

	return c.Redirect(tinu.URL, fiber.StatusTemporaryRedirect)
}

func updateTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var tinu model.Tinu
	err := c.BodyParser(&tinu)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	err = model.UpdateTinu(tinu)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"updated successfully",
		"data":tinu,
	})
}

func deleteTinu(c *fiber.Ctx) error {
	id := c.Params("id")
	err := model.DeleteTinu(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message":"deleted successfully",
		"id":id,
	})
}

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/:id", tinu)

	router.Get("/tinu", getAll)
	router.Get("/tinu/:id", getOne)
	router.Post("/tinu", createTinu)
	router.Patch("/tinu", updateTinu)
	router.Delete("/tinu/:id", deleteTinu)

	router.Listen(":3000")
}