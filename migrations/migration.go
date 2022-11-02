package migrations

import (
	"app-indihomesmart/infra/database"
	"app-indihomesmart/models"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{&models.Article{}, &models.User{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
