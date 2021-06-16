package check

import (
	"github.com/gin-gonic/gin"
	"github.com/gunjdesai/go-gin-boilerplate/handlers"
)

func Load(router *gin.RouterGroup) {

	check := router.Group("/check")

	check.GET("/status", handlers.Status)
	check.GET("/health", handlers.Health)

}
