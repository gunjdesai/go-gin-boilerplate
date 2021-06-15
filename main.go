package main

import (
	"api/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()

	routers.Load(server)

	server.Run(":8080")

}
