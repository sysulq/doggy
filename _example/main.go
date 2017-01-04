package main

import (
	"net/http"
	"time"

	"github.com/hnlq715/doggy"
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
	n.Use(doggy.NewPrometheus())
	n.UseHandler(m)

	doggy.ListenAndServe(n)
}
