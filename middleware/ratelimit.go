package middleware

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"
)

// rb is a token bucket that fills the bucket
// at the rate of rate tokens per second up to the given
// maximum capacity. Because of limited clock resolution,
// at high rates, the actual rate may be up to 1% different from the
// specified rate.
var rb *ratelimit.Bucket

type RateLimit struct {
	Rate     float64
	Capacity int64
}

// NewRateLimit returns a new RateLimit instance
func NewRateLimit(rate float64, capacity int64) *RateLimit {
	return &RateLimit{
		Rate:     rate,
		Capacity: capacity,
	}
}

// RateLimit is an Middleware that acts as a
// request throttler based on a token-bucket algorithm. Requests that would
// exceed the maximum request rate are delayed via the parameterized sleep
// function.
func (m *RateLimit) ServerHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if rb == nil {
		rb = ratelimit.NewBucketWithRate(m.Rate, m.Capacity)
	}

	time.Sleep(rb.Take(1))
	next(rw, r)
}
