package globals

import (
	"fmt"
	"os"

	"github.com/gunjdesai/go-gin-boilerplate/conf"
	"github.com/gunjdesai/go-gin-boilerplate/helpers"
	"github.com/gunjdesai/go-gin-boilerplate/loggers"
)

var (
	Config *conf.Configuration
	Log    *loggers.Logger
)

func Bootstrap() {

	var err error
	var confFile = helpers.GetEnv() + dotSeperator + fileType
	Config, err = conf.New(configPath, confFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(10)
	}

	Log, err = loggers.New(Config.App.Log.Level)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(10)
	}

	if err != nil {
		Log.Warn(err.Error())
		os.Exit(10)
	}

}
