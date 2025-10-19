package routes

import (
	"video-chat/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) *gin.RouterGroup {
	apiGroup := router.Group("/")
	apiGroup.GET("/health", handlers.HealthHandler)
	return apiGroup
}
