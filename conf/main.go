package conf

import (
	"errors"
	"path"

	env "github.com/Netflix/go-env"
	"github.com/gunjdesai/go-gin-boilerplate/helpers"
	"github.com/spf13/viper"
)

type Configuration struct {
	App struct {
		Host string
		Port string
		Log  struct {
			Level string `env:"APP_LOG_LEVEL,required=true"`
		}
	}
	Db struct {
		Elastic struct {
			Host string
			Port string
		}
	}
}

var Config *Configuration

func New(location string, fileName string) (*Configuration, error) {

	Config = &Configuration{}
	filePath := path.Join(helpers.GetRootDir(), location, fileName)
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		err = errors.New("Error reading config file, " + err.Error())
		return nil, err
	}

	if err := viper.Unmarshal(Config); err != nil {
		err = errors.New("Unable to decode to struct, " + err.Error())
		return nil, err
	}

	if _, err := env.UnmarshalFromEnviron(Config); err != nil {
		err = errors.New("Envirnoment variable override failure, " + err.Error())
		return nil, err
	}

	return Config, nil

}
