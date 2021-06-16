package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gunjdesai/go-gin-boilerplate/middlewares"
	"github.com/gunjdesai/go-gin-boilerplate/routers/check"
	v1 "github.com/gunjdesai/go-gin-boilerplate/routers/v1"
)

func Load(router *gin.Engine) {

	router.Use(middlewares.SetContentType)
	base := router.Group("/")
	api := router.Group("/api")

	check.Load(base)
	v1.Load(api)

}
