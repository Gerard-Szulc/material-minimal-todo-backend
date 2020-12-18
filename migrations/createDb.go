package migrations

import (
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateDb() {
	todo := &models.Todo{}
	database.DB.AutoMigrate(&todo)
}
