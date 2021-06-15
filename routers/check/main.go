package check

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.RouterGroup) {

	check := router.Group("/check")

	check.GET("/status", handlers.Status)
	check.GET("/health", handlers.Health)

}
