package controllers

import (
	"fmt"
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetTodos(c *fiber.Ctx) error {

	todos := &[]models.Todo{}

	database.DB.Select("id, title, completed").Find(&todos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": map[string]*[]models.Todo{
			"todos": todos,
		},
	})
}

func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
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

	todo := &models.Todo{
		Title:     body.Title,
		Completed: false,
	}
	database.DB.NewRecord(todo)

	database.DB.Create(&todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": map[string]*models.Todo{
			"todo": todo,
		},
	})
}

func GetTodo(c *fiber.Ctx) error {
	todo := &models.Todo{}

	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
		})
	}

	if database.DB.Where("id = ? ", id).First(&todo).RecordNotFound() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": map[string]*models.Todo{
			"todo": todo,
		},
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	type Request struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	todo := &models.Todo{}
	if database.DB.Where("id = ? ", id).First(&todo).RecordNotFound() {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo not found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}
	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	database.DB.Save(&todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": map[string]*models.Todo{
			"todo": todo,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	database.DB.Delete(&models.Todo{}, id)

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
