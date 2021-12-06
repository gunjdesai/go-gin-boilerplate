package main

import (
	"github.com/gunjdesai/go-gin-boilerplate/app"
	"github.com/gunjdesai/go-gin-boilerplate/globals"
	_ "go.uber.org/automaxprocs"
)

func main() {

	// Load Global variables
	globals.Bootstrap()

	// Load Routers and start app
	app.Start()

}
