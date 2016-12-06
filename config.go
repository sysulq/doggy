package doggy

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
	"github.com/uber-go/zap"
)

type Config struct {
	Listen     string           `ini:"listen"`
	Env        string           `ini:"env"`
	Logger     LoggerConfig     `ini:"log"`
	Middleware MiddlewareConfig `ini:"middleware"`
}

type LoggerConfig struct {
	File  *os.File  `ini:"-"`
	Level zap.Level `ini:"level"`
	Dir   string    `ini:"dir"`
}

type MiddlewareConfig struct {
	Timeout time.Duration `ini:"timeout"`
}

var configPath = "config.ini"

func SetConfigPath(path string) {
	configPath = path
}

var config Config

func LoadConfig(v, source interface{}, others ...interface{}) error {
	err := ini.MapTo(v, source, others...)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	config = Config{
		Listen: "0.0.0.0:8000",
		Env:    "dev",
		Logger: LoggerConfig{
			Level: zap.DebugLevel,
			File:  os.Stdout,
		},
		Middleware: MiddlewareConfig{
			Timeout: 5 * time.Second,
		},
	}

	if err := LoadConfig(&config, configPath); err != nil {
		log.Panic("LoadConfig failed", err)
	}

	if config.Env == "prod" {
		config.Logger.Level = zap.ErrorLevel
	}

	if len(config.Logger.Dir) != 0 {
		name, err := filepath.Abs(config.Logger.Dir)
		if err != nil {
			log.Panic("os.OpenFile failed", err)
		}
		l, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			log.Panic("os.OpenFile failed", err)
		}
		config.Logger.File = l
	}
}
