package doggy

import (
	"time"

	"github.com/spf13/viper"
)

func loadConfig() error {
	viper.SetDefault("listen", "0.0.0.0:8000")
	viper.SetDefault("env", "dev")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.file", "")

	timeout, _ := time.ParseDuration("5s")
	viper.SetDefault("middleware.timeout", timeout)
	viper.SetDefault("middleware.ratelimit", 5000)
	viper.SetDefault("middleware.capacity", 5000)

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
