package doggy

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/hnlq715/doggy/middleware"
	"github.com/hnlq715/xgen/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

// New returns a new Negroni instance with no middleware preconfigured.
func New(handlers ...negroni.Handler) *negroni.Negroni {
	return negroni.New()
}

// Classic returns a new Negroni instance with the default middleware already
// in the stack.
//
// Recovery - Panic Recovery Middleware
// Logger - Request/Response Logging
// TraceID - Trace ID Middleware
// RealIP - Get Real Client IP
// CloseNotify - Notify Client Close
// Timeout - Stop Process When Timeout
func Classic() *negroni.Negroni {
	n := negroni.New()
	n.Use(middleware.NewRecovery())
	logViper := viper.Sub("log")
	n.Use(middleware.NewLogger(logViper.GetString("level"), logViper.GetString("file")))
	n.Use(middleware.NewTraceID())
	n.Use(middleware.NewRealIP())
	n.Use(middleware.NewCloseNotify())
	n.Use(middleware.NewTimeout(viper.GetDuration("middleware.timeout")))
	return n
}

// NewMux returns a new router instance.
func NewMux() *mux.Router {
	return mux.NewRouter()
}

// NewHttpRouter returns a new httprouter instance.
func NewHttpRouter() *httprouter.Router {
	return httprouter.New()
}

// ListenAndServe always returns a non-nil error.
func ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(viper.GetString("listen"), handler)
}

// ListenAndServeGracefully always returns a non-nil error.
func ListenAndServeGracefully(handler http.Handler) error {
	l := utils.LogFromContext(context.Background())

	h := &http.Server{Addr: viper.GetString("listen"), Handler: handler}
	go func() {
		if err := h.ListenAndServe(); err != nil {
			l.Error("h.ListenAndServe failed", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	l.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	h.Shutdown(ctx)

	l.Info("Server gracefully stopped")
	return nil
}

func init() {
	// Load default config
	err := loadConfig()
	if err != nil {
		fmt.Printf("loadConfig failed, err=%s\n", err)
	}
}
