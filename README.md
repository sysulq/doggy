Doggy
===
[![Build Status](https://travis-ci.org/hnlq715/doggy.svg?branch=master)](https://travis-ci.org/hnlq715/doggy)
[![Go Report Card](https://goreportcard.com/badge/github.com/hnlq715/doggy)](https://goreportcard.com/report/github.com/hnlq715/doggy)

Lightweight, idiomatic and stable for building Go 1.7+ HTTP services.
It aims to provide a composable way to develop HTTP services.

dependency
---

* [uber-go/zap](github.com/uber-go/zap)
* [gorilla/mux](github.com/gorilla/mux)
* [gorilla/schema](github.com/gorilla/schema)
* [urfave/negroni](github.com/urfave/negroni)
* [juju/ratelimit](github.com/juju/ratelimit)
* [unrolled/render](github.com/unrolled/render)
* [asaskevich/govalidator.v4](gopkg.in/asaskevich/govalidator.v4)
* [julienschmidt/httprouter](github.com/julienschmidt/httprouter)

Generate api struct
---
* [himeraCoder/gojson](https://github.com/ChimeraCoder/gojson)
```
curl -s https://api.github.com/repos/chimeracoder/gojson | gojson -name=Repository -tags=schema,json
```

Generate model package
---
* [knq/xo](https://github.com/knq/xo)
```
xo mysql://user:passwd@host:port/db -o . --template-path templates --ignore-fields updateTime
```

Example
---
```
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
```
