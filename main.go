package main

import (
	"fmt"
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/migrations"
	"github.com/Gerard-Szulc/material-minimal-todo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func start () {
	database.InitDatabase()
	startApi()
}

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")
	routes.TodoRoute(api.Group("/todos"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "You are at the endpoint 😉",
		})
	})

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})
}
func startApi () {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		fmt.Println(exists)
	}
	fmt.Println("App is working on port :" + string(port))

	err := app.Listen(":"+string(port))

	if err != nil {
		panic(err)
	}
}


func main() {
	argsWithProg := os.Args
	if len(argsWithProg) <= 1 {
		start()
	} else {
		switch argsWithProg[1] {
		case "migrate":
			{
				database.InitDatabase()
				migrations.CreateDb()
				println("migration successful")

				return
			}
		case "start":
			{
				start()
			}
		}
	}
}
