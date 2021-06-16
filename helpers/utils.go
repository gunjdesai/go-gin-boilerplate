package helpers

import "os"

func GetEnv() string {

	var env string

	switch os.Getenv(ENV_VAR) {

	case DEVELOPMENT:
		env = DEVELOPMENT
	case STAGING:
		env = STAGING
	case PRODUCTION:
		env = PRODUCTION
	default:
		env = PRODUCTION

	}

	return env

}
