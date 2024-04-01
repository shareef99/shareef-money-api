package example

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	// userGroup := api.Group("/example")

	api.GET("/example", GetExample)
	api.POST("/example", CreateExample)
}
