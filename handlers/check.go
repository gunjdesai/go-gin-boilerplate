package handlers

import (
	"github.com/gin-gonic/gin"
)

func Status(context *gin.Context) {

	context.JSON(200, gin.H{"status": "up"})

}

func Health(context *gin.Context) {

	context.JSON(200, gin.H{"status": "up"})

}
