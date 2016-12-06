package doggy

import (
	"net/http"

	"github.com/uber-go/zap"
)

func Recovery(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log := LogFromContext(r.Context())
			log.Error("Panic", zap.Stack())
		}
	}()

	next(rw, r)
}
