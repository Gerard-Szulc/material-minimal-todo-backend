package controllers

import (
	"fmt"
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	"github.com/Gerard-Szulc/material-minimal-todo/utils"
	"github.com/gofiber/fiber/v2"
	utils2 "github.com/gofiber/fiber/v2/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const ExpirationMinutes = 10

func Login(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Pass     string `json:"pass"`
	}

	var body Request

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	valid := utils.Validation(
		[]models.Validation{
			{Value: body.Username, Valid: "username"},
			{Value: body.Pass, Valid: "password"},
		})
	if valid {
		user := &models.User{}
		if database.DB.Where("username = ? ", body.Username).First(&user).RecordNotFound() {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User not found",
			})
		}
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "Wrong password",
			})
		}
		if !user.Active {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "Account not active",
			})
		}

		tokenExpiryTimeUnix := time.Now().Add(time.Minute * ExpirationMinutes).Unix()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": utils.PrepareToken(user, tokenExpiryTimeUnix),
		})

	} else {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Values are not valid",
		})
	}
}

func Register(c *fiber.Ctx) error {

	type Request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Pass     string `json:"pass"`
	}

	var body Request

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	valid := utils.Validation(
		[]models.Validation{
			{Value: body.Username, Valid: "username"},
			{Value: body.Email, Valid: "email"},
			{Value: body.Pass, Valid: "password"},
		})

	if valid {
		generatedPassword := utils.HashAndSalt([]byte(body.Pass))
		user := &models.User{Username: body.Username, Email: body.Email, Password: generatedPassword, Uuid: utils2.UUIDv4()}
		database.DB.Create(&user)

		tokenExpiryTimeUnix := time.Now().Add(time.Minute * ExpirationMinutes).Unix()

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": utils.PrepareToken(user, tokenExpiryTimeUnix),
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Values are not valid",
		})
	}

}
