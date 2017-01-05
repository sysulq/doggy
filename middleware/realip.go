package middleware

import (
	"net/http"
	"strings"
)

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

type RealIP struct {
}

// NewRealIP returns a new RealIP instance
func NewRealIP() *RealIP {
	return &RealIP{}
}

func (m *RealIP) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if rip := realIP(r); rip != "" {
		r.RemoteAddr = rip
	}

	next(rw, r)
}

func realIP(r *http.Request) string {
	var ip string

	if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	} else if xrip := r.Header.Get(xRealIP); xrip != "" {
		ip = xrip
	}

	return ip
}
