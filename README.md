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
