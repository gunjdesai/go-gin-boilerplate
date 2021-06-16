package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gunjdesai/go-gin-boilerplate/handlers"
)

func Load(router *gin.RouterGroup) {

	v1 := router.Group("/v1")

	v1.GET("/sample-api", handlers.SampleApi)

}
