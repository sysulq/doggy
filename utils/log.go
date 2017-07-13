package utils

import (
	"context"
	"os"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logKey = "logKey"

func LogFromContext(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(logKey).(*zap.Logger)
	if !ok {
		env := viper.GetString("env")
		conf := zap.NewDevelopmentEncoderConfig()
		if env == "prod" {
			conf = zap.NewProductionEncoderConfig()
		}
		enc := zapcore.NewJSONEncoder(conf)
		return zap.New(
			zapcore.NewCore(enc,
				newWriteSyncer(),
				zapcore.Level(viper.GetInt("log.level"))),
			zap.AddCaller())
	}
	return l
}

func ContextWithLog(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey, log)
}

func newWriteSyncer() zapcore.WriteSyncer {
	if len(viper.GetString("log.file")) == 0 {
		return zapcore.AddSync(os.Stdout)
	}

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log.file"),
		MaxSize:    viper.GetInt("log.maxsize"), //megabytes
		MaxAge:     viper.GetInt("log.maxage"),  //days
		MaxBackups: viper.GetInt("log.backups"), //files
	})
}
