package main

import (
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add your frontend URL here
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	apiV1 := router.Group("/api/v1")

	users.RegisterRoutes(apiV1)
	example.RegisterRoutes(apiV1)

	router.Run()
}
