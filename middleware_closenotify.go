package doggy

import (
	"context"
	"net/http"
)

const StatusClientClosedRequest = 499

func CloseNotify(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	if cn, ok := rw.(http.CloseNotifier); ok {
		closeNotify := cn.CloseNotify()
		done := ctx.Done()

		go func() {
			select {
			case <-done:
				return
			case <-closeNotify:
				cancel()
				rw.WriteHeader(StatusClientClosedRequest)
				return
			}
		}()
	}

	next(rw, r.WithContext(ctx))
}
