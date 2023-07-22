package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/joho/godotenv"
	model "github.com/kehiy/tinu/models"
	"github.com/kehiy/tinu/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if !utils.ValidateEmail(user.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid email",
		})
	}

	if len(user.PassWord) < 7 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "enter a strong password",
		})
	}

	passWord, _ := bcrypt.GenerateFromPassword([]byte(user.PassWord), 10)
	user.PassWord = string(passWord)

	user.ID, err = utils.GenerateId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	err = model.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "OK",
		"data":    user,
	})
}

func UserLogin(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	dbUser, err := model.GetOneUser(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid password or email",
		})
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.PassWord), 10)
	if string(hashedPass) == dbUser.PassWord {
		// _ = godotenv.Load()
		secret := os.Getenv("JWTSECRET")
		secretKey := []byte(secret)
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = dbUser.ID                              // Subject
		claims["email"] = dbUser.Email                        // Name
		claims["iat"] = time.Now().Unix()                     // Issued At (current time)
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expiration Time (24 hours from now)

		signedToken, err := token.SignedString(secretKey)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "internal server error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": signedToken,
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "invalid password or email",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	authHeader := c.Get("authorization")
	secret := os.Getenv("JWTSECRET")
	parsedToken, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid password or email",
		})
	}

	if !parsedToken.Valid {
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)

	error := model.DeleteUser(userID)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "user was deleted successfully",
	})
}
