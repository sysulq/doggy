package doggy

import (
	"net/http"

	"github.com/uber-go/zap"
)

type Recovery struct {
}

// NewRecovery returns a new Recovery instance
func NewRecovery() *Recovery {
	return &Recovery{}
}

// Recovery is a doggy middleware that recovers from any panics and writes a 500.
func (m *Recovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log := LogFromContext(r.Context())
			log.Error("Panic", zap.Stack())
		}
	}()

	next(rw, r)
}
