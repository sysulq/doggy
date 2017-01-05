package doggy

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hnlq715/doggy/middleware"
	"github.com/julienschmidt/httprouter"
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
	n.Use(middleware.NewLogger(config.Logger.Level, config.Logger.File))
	n.Use(middleware.NewTraceID())
	n.Use(middleware.NewRealIP())
	n.Use(middleware.NewCloseNotify())
	n.Use(middleware.NewTimeout(config.Middleware.Timeout))
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
	return http.ListenAndServe(config.Listen, handler)
}

func init() {
	// Load default config
	err := initConfig()
	if err != nil {
		log.Panic(err)
	}

	initHttp()
}
