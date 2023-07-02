package controllers

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/kehiy/tinu/models"
	"github.com/kehiy/tinu/utils"
)

func CreateTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var tinu model.Tinu
	err := c.BodyParser(&tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	tinu.ID, err = utils.GenerateId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	err = model.CreateTinu(tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "OK",
		"data":    tinu,
	})
}

func Tinu(c *fiber.Ctx) error {
	id := c.Params("id")
	tinu, err := model.GetOneTinu(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	tinu.Clicked += 1
	err = model.UpdateTinu(tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Redirect(tinu.URL, fiber.StatusMovedPermanently)
}

func UpdateTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var tinu model.Tinu
	err := c.BodyParser(&tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	err = model.UpdateTinu(tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "updated successfully",
		"data":    tinu,
	})
}

func DeleteTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var tinu model.Tinu

	err := c.BodyParser(&tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	err = model.DeleteTinu(tinu.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error " + err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "deleted successfully",
		"id":      tinu.ID,
	})
}
