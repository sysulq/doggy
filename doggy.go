package doggy

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func New(handlers ...negroni.Handler) *negroni.Negroni {
	return negroni.New()
}

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

func NewMux() *mux.Router {
	return mux.NewRouter()
}

func init() {
	LoadConfig("config.ini")
}
