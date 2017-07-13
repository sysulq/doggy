package middleware

import (
	"net/http"

	"github.com/hnlq715/doggy/utils"
	"go.uber.org/zap"
)

type ParseForm struct {
}

// NewParseForm returns a new Recovery instance
func NewParseForm() *ParseForm {
	return &ParseForm{}
}

// ParseForm is a doggy middleware that call ParseForm for every http request.
func (m *ParseForm) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := r.ParseForm()
	if err != nil {
		utils.LogFromContext(r.Context()).Error("r.ParseForm failed", zap.Error(err))
	}

	next(rw, r)
}
