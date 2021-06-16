package handlers

import "github.com/gin-gonic/gin"

func SampleApi(context *gin.Context) {

	context.JSON(200, gin.H{"status": "working"})

}
