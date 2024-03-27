package initializers

import "github.com/shareef99/shareef-money-api/features/users"

func SyncDatabase() {
	DB.AutoMigrate(&users.User{})
}
