package migrations

import (
	"github.com/Gerard-Szulc/material-minimal-todo/database"
	"github.com/Gerard-Szulc/material-minimal-todo/models"
	"time"
)

func Migration191220201() {
	migration := &models.Migration{
		Version: "191220201",
		Date:    time.Now(),
	}

	if database.DB.Where("version = ? ", migration.Version).First(&migration).RecordNotFound() {
		database.DB.AutoMigrate(&migration)
		database.DB.NewRecord(migration)
		database.DB.Create(&migration)
	}

}
