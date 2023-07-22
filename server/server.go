package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kehiy/tinu/controllers"
	"github.com/kehiy/tinu/middlewares"
)

func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// tinu
	router.Get("/:id", controllers.Tinu)
	router.Get("/c/:id", controllers.CheckTinu)
	router.Post("/tinu", middlewares.Authenticate, controllers.CreateTinu)
	router.Patch("/tinu", middlewares.Authenticate, controllers.UpdateTinu)
	router.Delete("/tinu", middlewares.Authenticate, controllers.DeleteTinu)

	// user
	router.Post("/login", controllers.UserLogin)
	router.Post("/signup", controllers.CreateUser)
	router.Delete("/user/delete", controllers.DeleteUser)

	log.Fatal(router.Listen(":3000"))
}
