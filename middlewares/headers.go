package middlewares

import "github.com/gin-gonic/gin"

func SetContentType(context *gin.Context) {

	context.Writer.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
	context.Next()

}
