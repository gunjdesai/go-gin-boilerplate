package routers

import (
	"api/routers/check"
	v1 "api/routers/v1"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {

	base := router.Group("/")
	api := router.Group("/api")

	check.Load(base)
	v1.Load(api)

}
