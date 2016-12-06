package doggy

import "net/http"

func ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(config.Listen, handler)
}
