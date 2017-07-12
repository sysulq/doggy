package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/uber-go/zap"
	"github.com/urfave/negroni"
)

const zapKey = "zapKey"

type Logger struct {
	Level string
	File  *os.File
}

var logLevel = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// NewLogger returns a new Logger instance
func NewLogger(level string, name string) *Logger {
	file := os.Stdout
	if len(name) > 0 {
		newFile, err := os.Open(name)
		if err != nil {
			fmt.Printf("os.Open %s failed, use os.Stdout, err=%s\n", name, err)
		} else {
			file = newFile
		}
	}
	return &Logger{
		Level: level,
		File:  file,
	}
}

func (m *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	now := time.Now()
	log := zap.New(zap.NewJSONEncoder(timeFormat("timestamp")), zap.AddCaller(), zap.Level(logLevel[m.Level]), zap.Output(m.File))
	ctx := ContextWithLog(r.Context(), log)
	ww := negroni.NewResponseWriter(rw)

	next(ww, r.WithContext(ctx))

	log.Info("Completed", zap.Float64("responsetime", time.Now().Sub(now).Seconds()),
		zap.String("path", r.URL.Path), zap.String("host", r.Host), zap.Int("code", ww.Status()))
}

func LogFromContext(ctx context.Context) zap.Logger {
	l, ok := ctx.Value(zapKey).(zap.Logger)
	if !ok {
		return zap.New(zap.NewJSONEncoder(timeFormat("timestamp")), zap.AddCaller())
	}
	return l
}

func ContextWithLog(ctx context.Context, log zap.Logger) context.Context {
	return context.WithValue(ctx, zapKey, log)
}

func timeFormat(key string) zap.TimeFormatter {
	return zap.TimeFormatter(func(t time.Time) zap.Field {
		return zap.String(key, t.Local().Format("2006-01-02T15:04:05.000Z0700"))
	})
}
