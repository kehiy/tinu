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

	// tinu
	router.Get("/:id", controllers.Tinu)
	router.Post("/tinu", controllers.CreateTinu)
	router.Patch("/tinu", controllers.UpdateTinu)
	router.Delete("/tinu", controllers.DeleteTinu)

	// user
	router.Post("/login", controllers.UserLogin)
	router.Post("/signup", controllers.CreateUser)
	router.Delete("/user/delete", controllers.DeleteUser)

	router.Listen(":3000")
}
