package migrations

import (
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	"time"
)

func Migration191220203() {
	migration := &models.Migration{
		Version: "191220203",
		Date:    time.Now(),
	}
	user := &models.User{}

	if database.DB.Where("version = ? ", migration.Version).First(&migration).RecordNotFound() {
		database.DB.AutoMigrate(&user)
		database.DB.NewRecord(migration)
		database.DB.Create(&migration)
	}

}
