package migrations

import (
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func CreateDb() {
	todo := &models.Todo{}
	migration := &models.Migration{
		Version: "181220201",
		Date:    time.Now(),
	}
	if database.DB.Where("version = ? ", migration.Version).First(&migration).RecordNotFound() {
		database.DB.AutoMigrate(&todo)
		database.DB.NewRecord(migration)
		database.DB.Create(&migration)
	}
}
