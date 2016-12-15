package doggy

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

// RateLimit is an Middleware that acts as a
// request throttler based on a token-bucket algorithm. Requests that would
// exceed the maximum request rate are delayed via the parameterized sleep
// function.
func RateLimit(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if rb == nil {
		rb = ratelimit.NewBucketWithRate(config.Middleware.Rate, config.Middleware.Capacity)
	}

	time.Sleep(rb.Take(1))
	next(rw, r)
}
