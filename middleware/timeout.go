package middleware

import (
	"context"
	"net/http"
	"time"
)

type Timeout struct {
	Timeout time.Duration
}

// NewTimeout returns a new Timeout instance
func NewTimeout(timeout time.Duration) *Timeout {
	return &Timeout{
		Timeout: timeout,
	}
}

func (m *Timeout) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx, cancel := context.WithTimeout(r.Context(), m.Timeout)
	defer func() {
		cancel()
		if ctx.Err() == context.DeadlineExceeded {
			rw.WriteHeader(http.StatusGatewayTimeout)
		}
	}()

	next(rw, r.WithContext(ctx))
}
