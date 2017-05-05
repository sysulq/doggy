package doggy

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
	"github.com/uber-go/zap"
)

type Config struct {
	Listen     string
	Env        string
	Logger     LoggerConfig
	Middleware MiddlewareConfig
	HttpClient HttpClientConfig
}

type LoggerConfig struct {
	File  *os.File
	Level int32
	Dir   string
}

type MiddlewareConfig struct {
	Timeout  time.Duration
	Rate     float64
	Capacity int64
}

type HttpClientConfig struct {
	Timeout time.Duration
	Retry   uint
}

var (
	config     *Config
	configFile = flag.String("c", "config.ini", "config file name")
)

func ConfigFile() string {
	return *configFile
}

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
	if err := ini.MapToWithMapper(config, ini.TitleUnderscore, name); err != nil {
		return err
	}

	if config.Env == "prod" {
		config.Logger.Level = int32(zap.ErrorLevel)
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
			Level: int32(zap.DebugLevel),
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

	return config.Load(*configFile)
}
