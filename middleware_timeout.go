package doggy

import (
	"context"
	"net/http"
)

func Timeout(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx, cancel := context.WithTimeout(r.Context(), config.Middleware.Timeout)
	defer func() {
		cancel()
		if ctx.Err() == context.DeadlineExceeded {
			rw.WriteHeader(http.StatusGatewayTimeout)
		}
	}()

	next(rw, r.WithContext(ctx))
}
