package intializers

import "main/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Account{}, &models.RefreshToken{}, &models.User{})

}
