package doggy

import (
	"github.com/gorilla/mux"
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
	n.UseFunc(Recovery)
	n.UseFunc(Logger)
	n.UseFunc(TraceID)
	n.UseFunc(RealIP)
	n.UseFunc(CloseNotify)
	n.UseFunc(Timeout)
	return n
}

// NewMux returns a new router instance.
func NewMux() *mux.Router {
	return mux.NewRouter()
}

func init() {
	// Load default config
	LoadConfig("config.ini")
}
