package doggy

import (
	"context"
	"net/http"
)

type Timeout struct {
}

// NewTimeout returns a new Timeout instance
func NewTimeout() *Timeout {
	return &Timeout{}
}

func (m *Timeout) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx, cancel := context.WithTimeout(r.Context(), config.Middleware.Timeout)
	defer func() {
		cancel()
		if ctx.Err() == context.DeadlineExceeded {
			rw.WriteHeader(http.StatusGatewayTimeout)
		}
	}()

	next(rw, r.WithContext(ctx))
}
