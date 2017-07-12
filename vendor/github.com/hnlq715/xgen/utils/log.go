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
		out := newWriteSyncer()
		enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		return zap.New(zapcore.NewCore(enc, out, zapcore.Level(viper.GetInt("log.level"))), zap.AddCaller())
	}
	return l
}

func ContextWithLog(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey, log)
}

func newWriteSyncer() zapcore.WriteSyncer {
	if len(viper.GetString("log.path")) == 0 {
		return zapcore.AddSync(os.Stdout)
	}

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log.path"),
		MaxSize:    viper.GetInt("log.maxsize"), //megabytes
		MaxAge:     viper.GetInt("log.maxage"),  //days
		MaxBackups: viper.GetInt("log.backups"), //files
	})
}
