package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gunjdesai/go-gin-boilerplate/conf"
	"github.com/gunjdesai/go-gin-boilerplate/logger"
	"github.com/gunjdesai/go-gin-boilerplate/routers"
	_ "go.uber.org/automaxprocs"
)

func main() {

	server := gin.New()

	routers.Load(server)

	fmt.Println("Log Level Set at", conf.Config.App.Log.Level)
	logger.Log.Info("Server Started on Port: " + conf.Config.App.Port)
	server.Run(":" + conf.Config.App.Port)

}
