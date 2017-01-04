package doggy

import (
	"log"

	"github.com/gorilla/mux"
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
	n.Use(NewRecovery())
	n.Use(NewLogger())
	n.Use(NewTraceID())
	n.Use(NewRealIP())
	n.Use(NewCloseNotify())
	n.Use(NewTimeout())
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

func init() {
	// Load default config
	err := initConfig()
	if err != nil {
		log.Panic(err)
	}

	initHttp()
}
