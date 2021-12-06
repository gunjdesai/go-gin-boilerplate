package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gunjdesai/go-gin-boilerplate/globals"
	"github.com/gunjdesai/go-gin-boilerplate/routers"
)

func Start() {

	// Init for gin framework
	server := gin.New()

	// Load Routes
	routers.Load(server)

	globals.Log.Info("Log Level: " + globals.Config.App.Log.Level)
	globals.Log.Info("Server Started on Port: " + globals.Config.App.Port)
	server.Run(":" + globals.Config.App.Port)

}
