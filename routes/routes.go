package routes

import (
	"github.com/gin-gonic/gin"
	"mutual-funds-api/handlers"
)

func InitializeRoutes() *gin.Engine {

	routes := gin.Default()

	routes.GET("/funds", handlers.GetFundsHandler)

	return routes
}
