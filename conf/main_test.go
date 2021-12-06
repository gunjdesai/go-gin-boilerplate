package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInit(t *testing.T) {

	os.Setenv("APP_LOG_LEVEL", "warn")
	_, fileTypeError := New("configs", "dev.txt")
	assert.NotNil(t, fileTypeError)

	_, doesNotExistError := New("configs", "random.yml")
	assert.NotNil(t, doesNotExistError)

	_, err := New("test_configs", "dev.yml")
	assert.NotNil(t, err)

	os.Unsetenv("APP_LOG_LEVEL")
	_, envErr := New("configs", "dev.yml")
	assert.NotNil(t, envErr)

	os.Setenv("APP_LOG_LEVEL", "warn")
	config, e := New("configs", "dev.yml")
	assert.Nil(t, e)
	assert.NotNil(t, config.App.Host)
	assert.NotNil(t, config.Db.Elastic.Host)

	cfg, er := New("configs", "prod.yml")
	assert.Nil(t, er)
	assert.NotNil(t, cfg.App.Host)
	assert.NotNil(t, cfg.Db.Elastic.Host)

}
