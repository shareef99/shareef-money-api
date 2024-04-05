package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/controllers"
)

func RegisterRoutes(api *gin.RouterGroup) {
	userGroup := api.Group("/users")
	userGroup.GET("/", controllers.GetUsers)
	userGroup.GET("/by-id", controllers.GetUser)
	userGroup.POST("/", controllers.CreateUser)
	userGroup.POST("/signin", controllers.Signin)
	userGroup.PUT("/", controllers.UpdateUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)

	categoryGroup := api.Group("/categories")
	categoryGroup.GET("/", controllers.GetCategories)
	categoryGroup.GET("/by-user", controllers.GetUserCategories)
	categoryGroup.POST("/", controllers.CreateCategory)
	categoryGroup.DELETE("/:id", controllers.DeleteCategory)
	categoryGroup.PUT("/", controllers.UpdateCategory)

	accountGroup := api.Group("/accounts")
	accountGroup.GET("/", controllers.GetAccounts)
	accountGroup.GET("/by-user", controllers.GetUserAccounts)
	accountGroup.POST("/", controllers.CreateAccount)
	accountGroup.PUT("/", controllers.UpdateAccount)
	accountGroup.DELETE("/:id", controllers.DeleteAccount)
}
