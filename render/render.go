package render

import (
	"net/http"

	"github.com/unrolled/render"
)

var stdRender = render.New()

// Data writes out the raw bytes as binary data.
func Data(w http.ResponseWriter, status int, v []byte) error {
	return stdRender.Data(w, status, v)
}

// HTML builds up the response from the specified template and bindings.
func HTML(w http.ResponseWriter, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error {
	return stdRender.HTML(w, status, name, binding, htmlOpt...)
}

// JSON marshals the given interface object and writes the JSON response.
func JSON(w http.ResponseWriter, status int, v interface{}) error {
	return stdRender.JSON(w, status, v)
}

// JSONP marshals the given interface object and writes the JSON response.
func JSONP(w http.ResponseWriter, status int, callback string, v interface{}) error {
	return stdRender.JSONP(w, status, callback, v)
}

// Text writes out a string as plain text.
func Text(w http.ResponseWriter, status int, v string) error {
	return stdRender.Text(w, status, v)
}

// XML marshals the given interface object and writes the XML response.
func XML(w http.ResponseWriter, status int, v interface{}) error {
	return stdRender.XML(w, status, v)
}
