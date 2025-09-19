package intializers

import "main/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Account{}) // this will create posts table if missing
}
