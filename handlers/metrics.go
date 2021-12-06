package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Prometheus() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}
