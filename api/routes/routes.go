package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shareef99/shareef-money-api/api/controllers"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/users", controllers.GetUsers)
	api.GET("/users/:id", controllers.GetUser)
	api.POST("/users", controllers.CreateUser)
	api.POST("/users/signin", controllers.Signin)
	api.PUT("/users", controllers.UpdateUser)
	api.DELETE("/users/:id", controllers.DeleteUser)

	api.GET("/accounts", controllers.GetAccounts)
	api.GET("/accounts/by-user/:id", controllers.GetUserAccounts)
	api.POST("/accounts", controllers.CreateAccount)
	api.PUT("/accounts", controllers.UpdateAccount)
	api.DELETE("/accounts/:id", controllers.DeleteAccount)

	api.GET("/categories", controllers.GetCategories)
	api.GET("/categories/:id", controllers.GetCategory)
	api.GET("/categories/by-user", controllers.GetUserCategories)
	api.POST("/categories", controllers.CreateCategory)
	api.PUT("/categories", controllers.UpdateCategory)
	api.DELETE("/categories/:id", controllers.DeleteCategory)

	api.GET("/sub-categories", controllers.GetSubCategories)
	api.GET("/sub-categories/:id", controllers.GetSubCategory)
	api.POST("/sub-categories", controllers.CreateSubCategory)
	api.PUT("/sub-categories", controllers.UpdateSubCategory)
	api.DELETE("/sub-categories/:id", controllers.DeleteSubCategory)

	api.GET("/transactions", controllers.GetTransactions)
	api.GET("/transactions/account/:id", controllers.GetAccountTransactions)
	api.GET("/transactions/user/:id", controllers.GetUserTransactions)
	api.GET("/transactions/category/:id", controllers.GetCategoryTransactions)
	api.POST("/transactions", controllers.CreateTransaction)
}
