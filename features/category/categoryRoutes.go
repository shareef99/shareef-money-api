package category

import "github.com/gin-gonic/gin"

func RegisterRouter(api *gin.RouterGroup) {
	categoryGroup := api.Group("/categories")

	categoryGroup.GET("/", GetCategories)
	categoryGroup.GET("/by-user", GetUserCategories)
	categoryGroup.POST("/", CreateCategory)
	categoryGroup.DELETE("/:id", DeleteCategory)
	categoryGroup.PUT("/", UpdateCategory)
}
