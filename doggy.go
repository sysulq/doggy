package doggy

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hnlq715/doggy/middleware"
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
	n.Use(middleware.NewParseForm())
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

func init() {
	// Load default config
	err := loadConfig()
	if err != nil {
		fmt.Printf("loadConfig failed, err=%s\n", err)
	}
}
