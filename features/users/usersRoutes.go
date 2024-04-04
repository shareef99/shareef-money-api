package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	// userGroup := api.Group("/user")

	api.GET("/users", GetUsers)
	api.GET("/users/by-id", GetUser)
	api.POST("/users", CreateUser)
	api.POST("/users/signin", Signin)
	api.PUT("/users", UpdateUser)
	api.DELETE("/users/:id", DeleteUser)
}
