package middleware

import (
	"net/http"
	"runtime"

	"github.com/hnlq715/doggy/utils"
	"go.uber.org/zap"
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
			stack := make([]byte, 1024*8)
			stack = stack[:runtime.Stack(stack, false)]

			log := utils.LogFromContext(r.Context())
			log.Error("Panic", zap.Stack(string(stack)))
		}
	}()

	next(rw, r)
}
