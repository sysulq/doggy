package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hnlq715/doggy/utils"
	"github.com/urfave/negroni"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	log := utils.LogFromContext(r.Context())
	ctx := utils.ContextWithLog(r.Context(), log)

	next(rw, r.WithContext(ctx))

	ww := negroni.NewResponseWriter(rw)
	log.Info("Completed", zap.Float64("responsetime", time.Now().Sub(now).Seconds()),
		zap.String("path", r.URL.Path), zap.String("host", r.Host), zap.Int("code", ww.Status()))
}
