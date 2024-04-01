package migrate

import (
	"github.com/shareef99/shareef-money-api/features/example"
	"github.com/shareef99/shareef-money-api/features/users"
	"github.com/shareef99/shareef-money-api/initializers"
)

func SyncDatabase() {
	initializers.DB.AutoMigrate(&users.User{}, &example.Example{})
}
