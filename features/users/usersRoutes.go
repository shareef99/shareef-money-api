package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	// userGroup := api.Group("/user")

	api.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User routes under v1!",
		})
	})
}
