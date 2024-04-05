package accounts

import "github.com/gin-gonic/gin"

func RegisterRouter(api *gin.RouterGroup) {
	accountGroup := api.Group("/accounts")

	accountGroup.GET("/", GetAccounts)
	accountGroup.GET("/by-user", GetUserAccounts)
	accountGroup.POST("/", CreateAccount)
	accountGroup.PUT("/", UpdateAccount)
	accountGroup.DELETE("/:id", DeleteAccount)
}
