package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/features/example"
	"github.com/shareef99/shareef-money-api/features/users"
	"github.com/shareef99/shareef-money-api/initializers"
	"github.com/shareef99/shareef-money-api/migrate"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
	migrate.SyncDatabase()
}

func main() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	users.RegisterRoutes(apiV1)
	example.RegisterRoutes(apiV1)

	router.Run()
}
