package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kehiy/tinu/controllers"
)

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/tinu", controllers.GetAllTinus)
	router.Get("/:id", controllers.Tinu)
	router.Get("/tinu/:id", controllers.GetOneTinu)
	router.Post("/tinu", controllers.CreateTinu)
	router.Patch("/tinu", controllers.UpdateTinu)
	router.Delete("/tinu", controllers.DeleteTinu)

	router.Listen(":3000")
}
