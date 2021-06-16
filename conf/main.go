package conf

import (
	"fmt"
	"os"

	env "github.com/Netflix/go-env"
	"github.com/gunjdesai/go-gin-boilerplate/helpers"
	"github.com/spf13/viper"
)

type Configuration struct {
	App struct {
		Host string
		Port string
		Log  struct {
			Level string `env:"APP_LOG_LEVEL"`
		}
	}
}

var Config *Configuration

func init() {

	Config = &Configuration{}

	environment := helpers.GetEnv()
	viper.SetConfigFile(CONFIG_PATH + environment + FILE_TYPE)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s... ", err)
		os.Exit(10)
	}

	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("Unable to decode into struct, %v... ", err)
		os.Exit(10)
	}

	if _, err := env.UnmarshalFromEnviron(Config); err != nil {
		fmt.Printf("Environment Variables override failure: , %v... ", err)
		os.Exit(10)
	}

}
