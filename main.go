package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/features/users"
	"github.com/shareef99/shareef-money-api/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	users.RegisterRoutes(apiV1)

	router.Run()
}
