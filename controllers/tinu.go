package controllers

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/kehiy/tinu/models"
	"github.com/kehiy/tinu/utils"
)

func CreateTinu(c *fiber.Ctx) error {
	c.Accepts("application/json")
	userID := c.Locals("userID").(string)

	var tinu model.Tinu
	err := c.BodyParser(&tinu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	tinu.UserID = userID

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
	userID := c.Locals("userID").(string)

	dbtinu, err := model.GetOneTinu(tinu.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if dbtinu.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "you have no access",
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
	userID := c.Locals("userID").(string)

	dbtinu, err := model.GetOneTinu(tinu.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	
	if dbtinu.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "you have no access",
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
