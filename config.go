package doggy

import (
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
	HttpClient HttpClientConfig `ini:"httpclient"`
}

type LoggerConfig struct {
	File  *os.File  `ini:"-"`
	Level zap.Level `ini:"level"`
	Dir   string    `ini:"dir"`
}

type MiddlewareConfig struct {
	Timeout  time.Duration `ini:"timeout"`
	Rate     float64       `ini:"rate"`
	Capacity int64         `ini:"capacity"`
}

type HttpClientConfig struct {
	Timeout time.Duration `ini:"timeout"`
	Retry   uint          `ini:"retry"`
}

var config *Config

// LoadSection loads and parses specific section from INI config file.
// It will return error if list contains nonexistent files.
func LoadSection(v interface{}, name string, section string) error {
	file, err := ini.Load(name)
	if err != nil {
		return err
	}

	return file.Section(section).MapTo(v)
}

// Load loads and parses INI config file.
// It will return error if list contains nonexistent files.
func (config *Config) Load(name string) error {
	if err := ini.MapTo(config, name); err != nil {
		return err
	}

	if config.Env == "prod" {
		config.Logger.Level = zap.ErrorLevel
	}

	if len(config.Logger.Dir) != 0 {
		name, err := filepath.Abs(config.Logger.Dir)
		if err != nil {
			return err
		}
		l, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			return err
		}
		config.Logger.File = l
	}
	return nil
}

func initConfig() error {
	config = &Config{
		Listen: "0.0.0.0:8000",
		Env:    "dev",
		Logger: LoggerConfig{
			Level: zap.DebugLevel,
			File:  os.Stdout,
		},
		Middleware: MiddlewareConfig{
			Timeout:  5 * time.Second,
			Rate:     5000,
			Capacity: 1000,
		},
		HttpClient: HttpClientConfig{
			Timeout: 5 * time.Second,
			Retry:   3,
		},
	}

	return config.Load("config.ini")
}
