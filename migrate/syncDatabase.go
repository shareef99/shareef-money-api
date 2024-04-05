package migrate

import (
	"github.com/shareef99/shareef-money-api/api/models"
	"github.com/shareef99/shareef-money-api/initializers"
)

func SyncDatabase() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Account{IsHidden: false},
		&models.Category{},
		&models.SubCategory{},
		&models.Transaction{Type: "expense"},
	)
}
