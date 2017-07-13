package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/hnlq715/doggy/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/negroni"
)

var (
	namespace = "doggy"

	requestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "http_request_count",
		Help:      "http request count.",
	}, []string{"code", "path"})

	requestLatencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "http_request_latency_histogram",
		Help:      "http request latency histogram.",
	}, []string{"path"})
)

type Prometheus struct {
}

// NewPrometheus returns a new Prometheus instance
func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

func (p *Prometheus) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	now := time.Now()

	ww := negroni.NewResponseWriter(rw)

	next(ww, r)

	elasped := time.Now().Sub(now).Seconds()
	requestCounter.WithLabelValues(strconv.Itoa(ww.Status()), r.URL.Path).Inc()
	requestLatencyHistogram.WithLabelValues(r.URL.Path).Observe(elasped)

	utils.LogFromContext(r.Context()).Info(r.URL.Path)
}

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestLatencyHistogram)
}
