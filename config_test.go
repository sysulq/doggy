package doggy

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	err := loadConfig()
	assert.NotNil(t, err)
	assert.Equal(t, "0.0.0.0:8000", viper.GetString("listen"))
	assert.Equal(t, "dev", viper.GetString("env"))
	assert.Equal(t, "debug", viper.GetString("log.level"))
	assert.Equal(t, 5*time.Second, viper.GetDuration("middleware.timeout"))
	assert.Equal(t, 5000, viper.GetInt("middleware.ratelimit"))
	assert.Equal(t, 5000, viper.GetInt("middleware.capacity"))
}
