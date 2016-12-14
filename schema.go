package doggy

import (
	"github.com/gorilla/schema"
)

// DecodeSchema decodes a map[string][]string to a struct.
func DecodeSchema(dst interface{}, src map[string][]string) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return decoder.Decode(dst, src)
}
