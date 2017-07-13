package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/hnlq715/doggy"
	"github.com/hnlq715/doggy/httpclient"
	"github.com/hnlq715/doggy/middleware"
	"github.com/hnlq715/doggy/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	m := doggy.NewMux()

	m.Handle("/metrics", promhttp.Handler())
	m.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		processTime := 4 * time.Second
		ctx := r.Context()
		select {
		case <-ctx.Done():
			return
		case <-time.After(processTime):
		}
		render.Text(w, 200, "pong")
	})

	m.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		u, _ := url.Parse("http://httpbin.org/get")
		u.RawQuery = r.Form.Encode()
		err := httpclient.Get(r.Context(), u.String()).ToJSON(&data)
		if err != nil {
			render.Text(w, 200, err.Error())
			return
		}
		render.JSON(w, 200, data)
	})

	n := doggy.Classic()
	n.Use(middleware.NewPrometheus())
	n.UseHandler(m)

	doggy.ListenAndServeGracefully(n)
}
