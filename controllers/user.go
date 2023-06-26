package controllers

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/kehiy/tinu/models"
	"github.com/kehiy/tinu/utils"
)

func GetAll(c *fiber.Ctx) error {
	tinus, err := model.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(tinus)
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	tinu, err := model.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": tinu,
	})
}

func CreateTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var tinu model.Tinu
	err := c.BodyParser(&tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	tinu.ID = utils.GenerateId()
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
	tinu, err := model.GetOne(id)
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

	return c.Redirect(tinu.URL, fiber.StatusPermanentRedirect)
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

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "deleted successfully",
		"id":      tinu.ID,
	})
}
