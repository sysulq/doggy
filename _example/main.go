package main

import (
	"doggy"
	"net/http"
	"time"
)

func main() {

	m := doggy.NewMux()

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

	n := doggy.NewStd()
	n.UseHandler(m)

	doggy.ListenAndServe(n)
}
