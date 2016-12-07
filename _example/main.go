package main

import (
	"doggy"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	m := doggy.NewMux()

	m.Handle("/metrics", promhttp.Handler())
	m.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		processTime := 2 * time.Second
		ctx := r.Context()
		select {
		case <-ctx.Done():
			return
		case <-time.After(processTime):
		}
		doggy.Text(w, 200, "pong")
	})

	n := doggy.Classic()
	n.UseFunc(doggy.Prometheus)
	n.UseHandler(m)

	doggy.ListenAndServe(n)
}
