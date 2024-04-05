package migrate

import (
	"github.com/shareef99/shareef-money-api/initializers"
	"github.com/shareef99/shareef-money-api/models"
)

func SyncDatabase() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Account{IsHidden: false}, &models.Category{}, &models.SubCategory{})
}
